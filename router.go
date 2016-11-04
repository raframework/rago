package rago

import (
	"github.com/raframework/rago/http"
)

type router struct {
	request     *http.Request
	response    *http.Response
	uriPatterns map[UriPattern]ResourceMethod
}

func newRouter(request *http.Request, response *http.Response, uriPatterns map[UriPattern]ResourceMethod) *router {
	return &router{
		request:     http.NewRequest(),
		response:    http.NewResponse(),
		uriPatterns: uriPatterns,
	}
}

func (r *router) match() {

}

func (r *router) callResourceAction() {

}
