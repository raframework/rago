package rago

import (
	"net/http"

	"github.com/raframework/rago/log"
)

var Logger = log.NewRaLogger()

func init() {
	Logger.WithLevel(log.LevelDebug)
}

type RequestHandler func(*Context)

type app struct {
	uriPatterns    map[UriPattern]ResourceMethod
	requestHandler RequestHandler
}

func NewApp(uriPatterns map[UriPattern]ResourceMethod) *app {
	Logger.Debug("rago: NewApp")
	return &app{
		uriPatterns: uriPatterns,
	}
}

func (a *app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	Logger.Debug("rago: app.ServeHTTP")
	context := NewContext(a.uriPatterns)

	a.requestHandler(context)
}

func (a *app) WithRequestHanlder(requestHandler RequestHandler) *app {
	a.requestHandler = requestHandler

	return a
}

func (a *app) Run(address string) {
	Logger.Debug("rago: app.Run")
	http.ListenAndServe(address, a)
}
