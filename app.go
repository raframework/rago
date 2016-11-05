package rago

import (
	"log"
	"net/http"
)

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[UriPattern]ResourceMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[UriPattern]ResourceMethod) *app {
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(a.uriPatterns)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a))
}
