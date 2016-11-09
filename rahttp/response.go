package rahttp

import (
	"net/http"
)

type Response struct {
	stdResponseWriter http.ResponseWriter
}

func NewResponse(stdResponseWriter http.ResponseWriter) *Response {
	return &Response{
		stdResponseWriter: stdResponseWriter,
	}
}

func (r *Response) withStatus(code int) *Response {
	return r
}

func (r *Response) Write(data string) {

}
