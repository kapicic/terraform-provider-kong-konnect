package httptransport

import (
	"bytes"
  "encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kong-sdk/internal/utils"
)

type Request struct {
	Method      string
	BaseUrl     string
	Path        string
	Headers     map[string]string
	PathParams  map[string]string
	QueryParams map[string]string
	Body        any
}

func NewRequest(method string, baseUrl string, path string, headers map[string]string, queryParams map[string]string) Request {
	return Request{
		Method:      method,
		BaseUrl:     baseUrl,
		Path:        path,
		Headers:     utils.CloneMap(headers),
		QueryParams: utils.CloneMap(queryParams),
		PathParams:  make(map[string]string),
	}
}


func (r *Request) Clone() Request {
	if r == nil {
		return Request{
			Headers:    		make(map[string]string),
			QueryParams:    make(map[string]string),
			PathParams: 		make(map[string]string),
		}
	}

	clone := *r
	clone.PathParams = utils.CloneMap(r.PathParams)
	clone.Headers = utils.CloneMap(r.Headers)

	// TODO: Clone body and query params
	return clone
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}

func (r *Request) GetBaseUrl() string {
	return r.BaseUrl
}

func (r *Request) SetBaseUrl(baseUrl string) {
	r.BaseUrl = baseUrl
}

func (r *Request) GetPath() string {
	return r.Path
}

func (r *Request) SetPath(path string) {
	r.Path = path
}

func (r *Request) GetHeader(header string) string {
	return r.Headers[header]
}

// Sets a header value, overwriting any existing value
func (r *Request) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *Request) GetPathParam(param string) string {
	return r.PathParams[param]
}

// Sets a path param value, overwriting any existing value
func (r *Request) SetPathParam(param string, value any) {
	r.PathParams[param] = fmt.Sprintf("%v", value)
}

func (r *Request) GetQueryParam(param string) string {
	return r.QueryParams[param]
}

// Sets a query param value, overwriting any existing value
func (r *Request) SetQueryParam(param string, value any) {
	r.QueryParams[param] = fmt.Sprintf("%v", value)
}

func (r *Request) GetBody() any {
	return r.Body
}

func (r *Request) SetBody(body any) {
	r.Body = body
}

// Adds a new header if it does not already exists
func (r *Request) AddHeader(paramName string, paramValue string) {
	_, hasValue := r.Headers[paramName]
	if !hasValue {
		r.Headers[paramName] = paramValue
	}
}

// Adds a new path param if it does not already exists
func (r *Request) AddPathParam(paramName string, paramValue any) {
	_, hasValue := r.PathParams[paramName]
	if !hasValue {
		r.PathParams[paramName] = fmt.Sprintf("%v", paramValue)
	}
}

// Adds a new query param if it does not already exists
func (r *Request) AddQueryParam(paramName string, paramValue any) {
	_, hasValue := r.QueryParams[paramName]
	if !hasValue {
		r.QueryParams[paramName] = fmt.Sprintf("%v", paramValue)
	}
}

func (r *Request) CreateHttpRequest() (*http.Request, error) {
	requestUrl := r.getRequestUrl()

	requestBody, err := r.bodyToBytesReader()
	if err != nil {
		return nil, err
	}

	var httpRequest *http.Request
	if requestBody == nil {
		httpRequest, err = http.NewRequest(r.Method, requestUrl, nil)
	} else {
		httpRequest, err = http.NewRequest(r.Method, requestUrl, requestBody)
	}

	for key, value := range r.Headers {
		httpRequest.Header.Set(key, fmt.Sprint(value))
	}

	return httpRequest, err
}

func (r *Request) getRequestUrl() string {
	requestPath := r.Path
	for paramName, paramValue := range r.PathParams {
		placeholder := "{" + paramName + "}"
		requestPath = strings.ReplaceAll(requestPath, placeholder, url.PathEscape(paramValue))
	}

	requestOptions := ""
	params := r.queryParamsToUrlValues()
	if len(params) > 0 {
		requestOptions = fmt.Sprintf("?%s", params.Encode())
	}

	return r.BaseUrl + requestPath + requestOptions
}

func (r *Request) bodyToBytesReader() (*bytes.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

  requestBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot convert r.Body to []byte")
	}

	reader := bytes.NewReader([]byte(requestBody))

	return reader, nil
}

func (r *Request) queryParamsToUrlValues() url.Values {
	params := url.Values{}
	for key, value := range r.QueryParams {
		params.Add(key, value)
	}

	return params
}
