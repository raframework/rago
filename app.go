package rago

import (
	"net/http"

	"github.com/raframework/rago/rahttp"
)

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[rahttp.UriPattern]rahttp.ResourceAndMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[rahttp.UriPattern]rahttp.ResourceAndMethod) *app {
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(a.uriPatterns, w, req)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	panic(http.ListenAndServe(address, a))
}
