package rago

import (
	"net/http"

	"github.com/raframework/rago/log"
)

func init() {
	log.SetLevel(log.LDebug)
}

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[UriPattern]ResourceMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[UriPattern]ResourceMethod) *app {
	log.Debug("rago: NewApp")
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Debug("rago: app.ServeHTTP")
	context := NewContext(a.uriPatterns)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	log.Debug("rago: app.Run")
	http.ListenAndServe(address, a)
}
