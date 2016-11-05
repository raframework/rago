package rago

import (
	"github.com/raframework/rago/http"
)

type Context struct {
	request  *http.Request
	response *http.Response
	router   *router
}

func NewContext(uriPatterns UriPatterns) *Context {
	return &Context{
		request:  http.NewRequest(),
		response: http.NewResponse(),
		router:   newRouter(request, response, uriPatterns),
	}
}

func (c *Context) MatchUriPattern() *Context {
	c.router.match()

	return c
}

func (c *Context) CallResourceAction() *Context {
	c.router.callResourceAction()

	return c
}

func (c *Context) Call() *Context {
	return c
}

func (c *Context) Respond() *Context {
	return c
}
