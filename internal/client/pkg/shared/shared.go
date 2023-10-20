package shared

type RequestOptions struct {
	Headers     map[string]string
	QueryParams map[string]string
}

func NewRequestOptions() RequestOptions {
	return RequestOptions{
		Headers:     make(map[string]string),
		QueryParams: make(map[string]string),
	}
}

func (opts *RequestOptions) AddHeader(headerName string, headerValue string) {
	opts.Headers[headerName] = headerValue
}

func (opts *RequestOptions) GetHeader(headerName string) string {
	return opts.Headers[headerName]
}

func (opts *RequestOptions) AddQueryParam(paramName string, paramValue string) {
	opts.QueryParams[paramName] = paramValue
}

func (opts *RequestOptions) GetQueryParam(paramName string) string {
	return opts.QueryParams[paramName]
}
