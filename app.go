package rago

import (
	"github.com/raframework/rago/http"
)

type app struct {
	request  *http.Request
	response *http.Response
	router   *router
}

func NewApp(uriPatterns map[UriPattern]ResourceMethod) *app {
	request := http.NewRequest()
	response := http.NewResponse()
	router := newRouter(request, response, uriPatterns)

	return &app{
		request:  request,
		response: response,
		router:   router,
	}
}

func (a *app) MatchUriPattern() *app {
	a.router.match()

	return a
}

func (a *app) CallResourceAction() *app {
	a.router.callResourceAction()

	return a
}

func (a *app) Call() *app {
	return a
}

func (a *app) Respond() *app {
	return a
}
