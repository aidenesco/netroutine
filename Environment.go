package netroutine

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var universalOptions []EnvironmentOption

type Environment struct {
	WorkingData map[string]interface{}
	ExportData  map[string]interface{}
	Status      error
	Logs        []string
	Responses   []*http.Response
	Client      *http.Client
	mu          sync.Mutex
	fullLogs    bytes.Buffer
	secrets     map[string]string
	maxRetry    int
	retrySleep  time.Duration
}

type Result struct {
	ExportData map[string]interface{}
	Status     string
	Logs       []string
}

func AddUniversalEnvironmentOption(option EnvironmentOption) {
	universalOptions = append(universalOptions, option)
}

type EnvironmentOption func(environment *Environment) error

func (wce *Environment) ToResult() *Result {
	return &Result{
		ExportData: wce.ExportData,
		Status:     wce.Status.Error(),
		Logs:       wce.Logs,
	}
}

func (wce *Environment) FullLogs() string {
	return wce.fullLogs.String()
}

func (wce *Environment) lastResponse() (*http.Response, error) {
	if len(wce.Responses) == 0 {
		return nil, errors.New("no responses to parse")
	}
	return wce.Responses[len(wce.Responses)-1], nil
}

func (wce *Environment) lastResponseBody() (string, error) {
	resp, err := wce.lastResponse()
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	//if string(respBody) == "" {
	//	return "", errors.New("empty body")
	//}

	return string(respBody), nil
}

func (wce *Environment) logHTTPResponse(resp *http.Response) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	wce.fullLogs.WriteString("[Request]\n")
	err = resp.Request.Write(&wce.fullLogs)
	if err != nil {
		return err
	}

	if resp.Request.Method != http.MethodGet {
		wce.fullLogs.WriteString("\n\n")
	}

	wce.fullLogs.WriteString("[Response]\n")
	err = resp.Write(&wce.fullLogs)
	if err != nil {
		return err
	}
	wce.fullLogs.WriteString("\n\n")

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))

	return nil
}

func (wce *Environment) addLog(message string) {
	_, _ = wce.fullLogs.Write([]byte(fmt.Sprintf("%v\n\n", message)))
	wce.Logs = append(wce.Logs, message)
}

func (wce *Environment) setData(toKey string, data interface{}) {
	wce.mu.Lock()
	defer wce.mu.Unlock()
	wce.WorkingData[toKey] = data
}

func (wce *Environment) setExportData(toKey string, data interface{}) {
	wce.mu.Lock()
	defer wce.mu.Unlock()
	wce.ExportData[toKey] = data
}

func (wce *Environment) getData(fromKey string) (data interface{}, found bool) {
	wce.mu.Lock()
	defer wce.mu.Unlock()
	data, found = wce.WorkingData[fromKey]
	return
}

func (wce *Environment) getSecret(fromKey string) (data string, found bool) {
	wce.mu.Lock()
	defer wce.mu.Unlock()
	data, found = wce.secrets[fromKey]
	return
}

func (wce *Environment) absorbData(data map[string]interface{}) {
	for k, v := range data {
		wce.WorkingData[k] = v
		wce.ExportData[k] = v
	}
}

func NewEnvironment(data map[string]interface{}, options ...EnvironmentOption) (env *Environment, err error) {
	env = newBaseEnvironment()
	env.absorbData(data)
	for _, f := range universalOptions {
		err = f(env)
		if err != nil {
			return
		}
	}

	for _, f := range options {
		err = f(env)
		if err != nil {
			return
		}
	}

	return
}

func WithUniqueTransport() EnvironmentOption {
	return func(environment *Environment) error {
		environment.Client.Transport = &http.Transport{
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          5,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		return nil
	}
}

func WithHTTPTimeout(duration time.Duration) EnvironmentOption {
	return func(environment *Environment) error {
		environment.Client.Timeout = duration
		return nil
	}
}

func WithProxyURL(proxy string) EnvironmentOption {
	return func(environment *Environment) error {
		purl, err := url.Parse(proxy)
		if err != nil {
			return err
		}

		trans, ok := environment.Client.Transport.(*http.Transport)
		if !ok {
			return errors.New("webchain: Client transport not *http.Transport")
		}

		trans.Proxy = http.ProxyURL(purl)
		return nil
	}

}

func WithProxyFunc(pf func(r *http.Request) (*url.URL, error)) EnvironmentOption {
	return func(environment *Environment) error {
		trans, ok := environment.Client.Transport.(*http.Transport)
		if !ok {
			return errors.New("webchain: Client transport not *http.Transport")
		}
		trans.Proxy = pf
		return nil
	}
}

func WithRetrySleep(duration time.Duration) EnvironmentOption {
	return func(environment *Environment) error {
		environment.retrySleep = duration
		return nil
	}
}

func WithRetryLimit(limit int) EnvironmentOption {
	return func(environment *Environment) error {
		environment.maxRetry = limit
		return nil
	}
}

func WithSecret(key, value string) EnvironmentOption {
	return func(environment *Environment) error {
		environment.secrets[key] = value
		return nil
	}
}

func newBaseEnvironment() *Environment {
	jar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Transport: http.DefaultTransport,
		Jar:       jar,
		Timeout:   time.Second * 15,
	}

	return &Environment{
		WorkingData: map[string]interface{}{},
		ExportData:  map[string]interface{}{},
		secrets:     map[string]string{},
		Status:      Success,
		Logs:        []string{},
		Client:      cli,
		maxRetry:    5,
		retrySleep:  time.Second,
	}
}

func toFloat64(data interface{}) (f float64, err error) {
	switch data.(type) {
	case string:
		f, err = strconv.ParseFloat(data.(string), 64)
	case int:
		f = float64(data.(int))
	case float64:
		f = data.(float64)
	default:
		err = fmt.Errorf("unable to convert value of type: %v", reflect.TypeOf(data))
	}
	return
}

func toString(data interface{}) (s string, err error) {
	switch data.(type) {
	case string:
		s = data.(string)
	case int:
		s = strconv.Itoa(data.(int))
	case float64:
		s = fmt.Sprintf("%v", data.(float64))
	default:
		s = fmt.Sprintf("%v", data)
	}
	return
}

func toTime(data interface{}) (time.Time, error) {
	ptime, ok := data.(time.Time)
	if !ok {
		return time.Time{}, errors.New("value not a time")
	}
	return ptime, nil
}

func statusFromString(s string) (retStatus error, retError error) {
	switch s {
	case Success.Error():
		retStatus = Success
	case Fail.Error():
		retStatus = Fail
	case Retry.Error():
		retStatus = Retry
	case Error.Error():
		retStatus = Error
	case Custom.Error():
		retStatus = Custom
	default:
		retError = errors.New("status should be one of: success, fail, retry, error, custom")
	}
	return
}

func log(b Runnable, m string, e error) (string, error) {
	return "[" + e.Error() + "] " + b.kind() + ": " + m, e
}
