package rago

import (
	"github.com/raframework/rago/http"
)

type Context struct {
	request  *http.Request
	response *http.Response
	router   *router
}

func NewContext(uriPatterns map[UriPattern]ResourceMethod) *Context {
	Logger.Debug("rago: NewContext")

	request := http.NewRequest()
	response := http.NewResponse()

	return &Context{
		request:  request,
		response: response,
		router:   newRouter(request, response, uriPatterns),
	}
}

func (c *Context) MatchUriPattern() *Context {
	Logger.Debug("rago: context.MatchUriPatter")

	c.router.match()

	return c
}

func (c *Context) CallResourceAction() *Context {
	Logger.Debug("rago: context.callResourceAction")

	c.router.callResourceAction()

	return c
}

func (c *Context) Call() *Context {
	Logger.Debug("rago: context.Call")

	return c
}

func (c *Context) Respond() *Context {
	Logger.Debug("rago: context.Respond")

	return c
}
