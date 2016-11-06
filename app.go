package rago

import (
	"net/http"

	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/ralog"
)

func init() {
	ralog.SetLevel(ralog.LDebug)
}

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[rahttp.UriPattern]rahttp.ResourceMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[rahttp.UriPattern]rahttp.ResourceMethod) *app {
	ralog.Debug("rago: NewApp")
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ralog.Debug("rago: app.ServeHTTP")
	context := NewContext(a.uriPatterns, w, req)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	ralog.Debug("rago: app.Run")

	ralog.Critical(http.ListenAndServe(address, a))
}
