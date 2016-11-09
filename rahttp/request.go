package rahttp

import (
	"net/http"

	"github.com/raframework/rago/ralog"
)

type Request struct {
	stdRequest        *http.Request
	matchedUriPattern UriPattern
	attributes        map[string]string
}

func NewRequest(stdRequest *http.Request) *Request {
	return &Request{
		stdRequest: stdRequest,
	}
}

func (r *Request) GetUriPath() string {
	path := r.stdRequest.URL.Path
	ralog.Debug("rahttp: uri path ", path)

	return path
}

func (r *Request) WithMatchedUriPattern(pattern UriPattern) {
	r.matchedUriPattern = pattern
}

func (r *Request) GetMethod() Method {
	return Method(r.stdRequest.Method)
}

func (r *Request) WithAttributes(attributes map[string]string) {
	r.attributes = attributes
}

func (r *Request) GetParsedBody() map[string]interface{} {

}

func (r *Request) GetQueryParams() map[string]string {

}

func (r *Request) GetAttributes() map[string]string {

}
