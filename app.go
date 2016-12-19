package rago

import (
	"net/http"

	"github.com/coderd/glog"
	"github.com/raframework/rago/rahttp"
)

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[rahttp.UriPattern]rahttp.ResourceAndMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[rahttp.UriPattern]rahttp.ResourceAndMethod) *app {
	glog.Debug("rago: NewApp")
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	glog.Debug("rago: app.ServeHTTP")
	context := NewContext(a.uriPatterns, w, req)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	glog.Debug("rago: app.Run")

	glog.Critical(http.ListenAndServe(address, a))
}
