package netroutine

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	blocks[idRequest] = &Request{}
}

const idRequest = "Request"

type Request struct {
	URL struct {
		URL       string
		Variables []string
		Complex   bool
	}
	BodyVar string
	Headers []struct {
		Key       string
		Value     string
		Variables []string
		Complex   bool
	}
	KeyChain []struct {
		Status     Status
		StatusCode int
		TextKey    string
	}
	Method         string
	IgnoreRedirect bool
}

func (b *Request) fromBytes(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Request) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Request) kind() string {
	return idRequest
}

func (b *Request) Run(ctx context.Context, wce *Environment) (string, Status) {
	builtURL, err := b.buildURL(wce)
	if err != nil {
		return log(b, reportError("building url", err), Error)
	}

	var (
		builtBody io.Reader
		resetBody io.Reader
	)

	if b.Method == http.MethodGet || b.Method == http.MethodHead || b.Method == http.MethodOptions || b.Method == "" {
		builtBody = bytes.NewBuffer([]byte{})
		resetBody = bytes.NewBuffer([]byte{})
	} else {
		v, ok := wce.getData(b.BodyVar)
		if !ok {
			return log(b, missingWorkingData(b.BodyVar), Error)
		}

		sv, err := toString(v)
		if err != nil {
			return log(b, reportWrongType(b.BodyVar), Error)
		}

		builtBody = strings.NewReader(sv)
		resetBody = strings.NewReader(sv)
	}

	req, err := http.NewRequestWithContext(ctx, b.Method, builtURL, builtBody)
	if err != nil {
		return log(b, reportError("building request", err), Error)
	}

	err = b.addHeaders(wce, req)
	if err != nil {
		return log(b, reportError("adding headers", err), Error)
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
		return log(b, reportError("doing request", err), Retry)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return log(b, reportError("reading request body", err), Error)
	}

	err = resp.Body.Close()
	if err != nil {
		return log(b, reportError("closing request body", err), Error)
	}

	wce.lastResponseBody = string(respBody)

	resp.Request.Body = ioutil.NopCloser(resetBody)
	resp.Body = ioutil.NopCloser(bytes.NewBufferString(wce.lastResponseBody))

	reqLogBuffer := new(bytes.Buffer)

	reqLogBuffer.WriteString("[Request]\n")
	err = resp.Request.Write(reqLogBuffer)
	if err != nil {
		return log(b, reportError("writing request logs", err), Error)
	}

	if resp.Request.ContentLength > 0 {
		reqLogBuffer.WriteString("\n")
	}

	reqLogBuffer.WriteString("[Response]\n")
	err = resp.Write(reqLogBuffer)
	if err != nil {
		return log(b, reportError("writing response logs", err), Error)
	}

	if resp.ContentLength > 0 {
		reqLogBuffer.WriteString("\n")
	}

	logs := base64.StdEncoding.EncodeToString(reqLogBuffer.Bytes())

	resp.Body = ioutil.NopCloser(bytes.NewBufferString(wce.lastResponseBody))
	wce.lastResponse = resp

	for _, key := range b.KeyChain {
		if (key.StatusCode == resp.StatusCode) && (strings.Contains(wce.lastResponseBody, key.TextKey)) {
			return log(b, fmt.Sprintf("found key: \"%s\" in %s", key.TextKey, logs), key.Status)
		}
	}

	return log(b, fmt.Sprintf("no keys found: %s", logs), Error)
}

func (b *Request) buildURL(wce *Environment) (string, error) {
	if !b.URL.Complex {
		return b.URL.URL, nil
	}

	var sub []interface{}
	for _, v := range b.URL.Variables {
		sv, ok := wce.getData(v)
		if !ok {
			return "", errors.New(missingWorkingData(v))
		}
		sub = append(sub, sv)
	}
	return fmt.Sprintf(b.URL.URL, sub...), nil
}

func (b *Request) addHeaders(wce *Environment, r *http.Request) error {
	for _, h := range b.Headers {
		if !h.Complex {
			r.Header.Add(h.Key, h.Value)
			continue
		}

		var sub []interface{}
		for _, v := range h.Variables {
			sv, ok := wce.getData(v)
			if !ok {
				return errors.New(missingWorkingData(v))
			}

			sub = append(sub, sv)
		}
		r.Header.Add(h.Key, fmt.Sprintf(h.Value, sub...))
	}
	return nil
}
