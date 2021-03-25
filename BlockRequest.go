package netroutine

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const idBlockRequest = "BlockRequest"

type BlockRequest struct {
	URL            URLBuilder
	BodyVar        string
	Headers        []HeaderBuilder
	KeyChain       []Key
	Method         string
	IgnoreRedirect bool
}

type URLBuilder struct {
	URL       string
	Variables []string
	Complex   bool
}

type HeaderBuilder struct {
	Key       string
	Value     string
	Variables []string
	Complex   bool
}

type Key struct {
	Flag       string
	StatusCode int
	TextKey    string
}

func (b *BlockRequest) fromBytes(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *BlockRequest) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockRequest) kind() string {
	return idBlockRequest
}

func (b *BlockRequest) Run(wce *Environment) (string, error) {
	builtURL, err := buildURL(b.URL, wce)
	if err != nil {
		return log(b, fmt.Sprintf("error building url - %v", err), Error)
	}

	var (
		builtBody io.Reader
		resetBody io.Reader
	)

	if b.Method == http.MethodGet || b.Method == "" {
		builtBody = bytes.NewBuffer([]byte{})
		resetBody = bytes.NewBuffer([]byte{})
	} else {
		bVar, ok := wce.getData(b.BodyVar)
		if !ok {
			return log(b, "unable to find body variable", Error)
		}

		sbVar, err := toString(bVar)
		if err != nil {
			return log(b, "unable to convert body variable to string", Error)
		}

		builtBody = strings.NewReader(sbVar)
		resetBody = strings.NewReader(sbVar)
	}

	req, err := http.NewRequest(b.Method, builtURL, builtBody)
	if err != nil {
		return log(b, fmt.Sprintf("error building request - %v", err), Error)
	}

	err = addHeaders(b.Headers, wce, req)
	if err != nil {
		return log(b, fmt.Sprintf("error building headers - %v", err), Error)
	}

	if b.IgnoreRedirect {
		wce.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		defer func() {
			wce.Client.CheckRedirect = nil
		}()
	}

	resp, err := wce.Client.Do(req)
	if err != nil {
		return log(b, fmt.Sprintf("error doing request - %v", err), Retry)
	}

	wce.Responses = append(wce.Responses, resp)

	resp.Request.Body = ioutil.NopCloser(resetBody)

	err = wce.logHTTPResponse(resp)
	if err != nil {
		return log(b, fmt.Sprintf("error logging response body - %v", err), Error)
	}

	strbody, err := wce.lastResponseBody()
	if err != nil {
		return log(b, fmt.Sprintf("error reading response body - %v", err), Error)
	}

	for _, key := range b.KeyChain {
		if (key.StatusCode == resp.StatusCode) && (strings.Contains(strbody, key.TextKey)) {
			switch key.Flag {
			case "success":
				return log(b, fmt.Sprintf("found success key %v", key.TextKey), Success)
			case "fail":
				return log(b, fmt.Sprintf("found fail key %v", key.TextKey), Fail)
			case "retry":
				return log(b, fmt.Sprintf("found retry key %v", key.TextKey), Retry)
			case "error":
				return log(b, fmt.Sprintf("found error key %v", key.TextKey), Error)
			case "custom":
				return log(b, fmt.Sprintf("found custom key %v", key.TextKey), Custom)
			}
		}
	}

	return log(b, fmt.Sprintf("no keys found - %v, %v, [%v]", builtURL, resp.StatusCode, base64.StdEncoding.EncodeToString([]byte(strbody))), Error)
}

func buildURL(b URLBuilder, wce *Environment) (string, error) {
	if !b.Complex {
		return b.URL, nil
	}

	var sub []interface{}
	for _, v := range b.Variables {
		sv, ok := wce.getData(v)
		if !ok {
			return "", errors.New(fmt.Sprintf("failed to find \"%s\" variable", v))
		}
		sub = append(sub, sv)
	}
	return fmt.Sprintf(b.URL, sub...), nil
}

func addHeaders(b []HeaderBuilder, wce *Environment, r *http.Request) error {
	for _, v := range b {
		if !v.Complex {
			r.Header.Add(v.Key, v.Value)
			continue
		}

		var sub []interface{}
		for _, v := range v.Variables {
			sv, ok := wce.getData(v)
			if !ok {
				return errors.New(fmt.Sprintf("Failed to find \"%v\" variable.", v))
			}

			sub = append(sub, sv)
		}
		r.Header.Add(v.Key, fmt.Sprintf(v.Value, sub...))
	}
	return nil
}
