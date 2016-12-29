package rago

import (
	"log"
	"net/http"

	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/rahttp"
)

type ErrorHandler func(interface{}, *rahttp.Request, *rahttp.Response)

type Processor interface {
	Process(*rahttp.Request, *rahttp.Response)
}

type Context struct {
	request      *rahttp.Request
	response     *rahttp.Response
	router       *router
	err          interface{}
	errorHandler ErrorHandler
}

func NewContext(uriPatterns map[rahttp.UriPattern]rahttp.ResourceAndMethod, w http.ResponseWriter, req *http.Request) *Context {
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
		c.handleError(err)
	}
}

func (c *Context) MatchUriPattern() *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	c.router.match()

	return c
}

func (c *Context) CallResourceAction() *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	c.router.callResourceAction()

	return c
}

func (c *Context) Call(p Processor) *Context {
	if c.err != nil {
		return c
	}
	defer c.recover()

	p.Process(c.request, c.response)

	return c
}

func (c *Context) Respond() *Context {
	c.response.FlushHeaders()
	c.response.FlushStatus()
	c.response.FlushBody()

	return c
}

func (c *Context) WithErrorHandler(errorHandler ErrorHandler) {
	c.errorHandler = errorHandler
}

func (c *Context) handleError(err interface{}) {
	c.err = err
	if c.errorHandler != nil {
		defer func() {
			if err := recover(); err != nil {
				c.defaultHandleError(err)
			}
		}()

		c.errorHandler(err, c.request, c.response)
	} else {
		c.defaultHandleError(err)
	}
}

func (c *Context) defaultHandleError(err interface{}) {
	var statusCode int

	raError, ok := err.(*raerror.RaError)
	if ok {
		switch raError.Typ() {
		case raerror.TypBadRequest:
			statusCode = 400
		case raerror.TypNotFound:
			statusCode = 404
		case raerror.TypMethodNotAllowed:
			statusCode = 405
		case raerror.TypUnsupportedMediaType:
			statusCode = 415
		default:
			statusCode = 500
		}
	} else {
		statusCode = 500
	}

	if statusCode == 500 {
		log.Println("[Error] rago: ", err)
	}

	c.response.WithStatus(statusCode)
}
