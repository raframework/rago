package rago

import (
	"net/http"
)

type app struct {
	uriPatterns    UriPatterns
	requestHandler RequestHandler
}

func NewApp(uriPatterns UriPatterns) *app {
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(a.uriPatterns)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler func(*Context)) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	http.ListenAndServe(address, a)
}
