package rago

import (
	"net/http"

	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/ralog"
)

type Context struct {
	request      *rahttp.Request
	response     *rahttp.Response
	router       *router
	err          error
	errorHandler func(interface{}, *rahttp.Request, *rahttp.Response)
}

func NewContext(uriPatterns map[rahttp.UriPattern]rahttp.ResourceMethod, w http.ResponseWriter, req *http.Request) *Context {
	ralog.Debug("rago: NewContext")

	request := rahttp.NewRequest(req)
	response := rahttp.NewResponse(w)

	return &Context{
		request:  request,
		response: response,
		router:   newRouter(request, response, uriPatterns),
	}
}

func (c *Context) recover() {
	if err := recover(); err != nil {
		c.errorHandler(err, c.request, c.response)
	}
}

func (c *Context) MatchUriPattern() *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	ralog.Debug("rago: context.MatchUriPatter")

	c.router.match()

	return c
}

func (c *Context) CallResourceAction() *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	ralog.Debug("rago: context.callResourceAction")

	c.router.callResourceAction()

	return c
}

func (c *Context) Call() *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	ralog.Debug("rago: context.Call")

	return c
}

func (c *Context) Respond() *Context {
	ralog.Debug("rago: context.Respond")
	c.response.FlushStatus()
	c.response.FlushHeaders()
	c.response.FlushBody()

	return c
}

func (c *Context) WithErrorHandler(errorHandler func(interface{}, *rahttp.Request, *rahttp.Response)) {
	c.errorHandler = errorHandler
}
