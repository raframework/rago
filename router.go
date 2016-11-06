package rago

import (
	"github.com/raframework/rago/http"
	"github.com/raframework/rago/log"
)

type router struct {
	request     *http.Request
	response    *http.Response
	uriPatterns map[UriPattern]ResourceMethod
}

func newRouter(request *http.Request, response *http.Response, uriPatterns map[UriPattern]ResourceMethod) *router {
	log.Debug("rago: NewRouter")

	return &router{
		request:     request,
		response:    response,
		uriPatterns: uriPatterns,
	}
}

func (r *router) match() {
	log.Debug("rago: router.match")
}

func (r *router) callResourceAction() {
	log.Debug("rago: router.callResourceAction")
}
