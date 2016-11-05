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
	Logger.Debug("rago: NewRouter")

	return &router{
		request:     request,
		response:    response,
		uriPatterns: uriPatterns,
	}
}

func (r *router) match() {
	Logger.Debug("rago: router.match")
}

func (r *router) callResourceAction() {
	Logger.Debug("rago: router.callResourceAction")
}
