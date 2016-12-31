package handler

import (
	"log"
	"runtime"

	"github.com/raframework/rago/example/app/config/code"
	"github.com/raframework/rago/example/app/lib/apperror"
	"github.com/raframework/rago/example/app/lib/rsp"
	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/rahttp"
)

func ErrorHandler(err interface{}, request *rahttp.Request, response *rahttp.Response) {
	appError, ok := err.(*apperror.AppError)
	if ok {
		handleAppError(appError, request, response)
		return
	}

	raError, ok := err.(*raerror.RaError)
	if ok {
		handleRaError(raError, request, response)
		return
	}

	log.Println("example:", err, "\nruntime stack:\n", getRuntimeStack())
	response.WithStatus(500)
	response.Write(rsp.ErrorJson(code.InternalServerError, ""))
}

func handleAppError(appError *apperror.AppError, request *rahttp.Request, response *rahttp.Response) {
	statusCode := appError.Typ()
	c := appError.Code()
	message := appError.Message()

	response.WithStatus(statusCode)
	response.Write(rsp.ErrorJson(c, message))
}

func handleRaError(raError *raerror.RaError, request *rahttp.Request, response *rahttp.Response) {
	var c, statusCode int
	var message string

	switch raError.Typ() {
	case raerror.TypMethodNotAllowed:
		statusCode = 405
		c = code.MethodNotAllowed

	case raerror.TypNotFound:
		statusCode = 404
		c = code.ResourceNotFound

	case raerror.TypBadRequest:
		statusCode = 400
		c = code.BadBody
		message = raError.Message()

	case raerror.TypUnsupportedMediaType:
		statusCode = 415
		c = code.UnsupportedMediaType
		message = raError.Message()

	default:
		statusCode = 500
		c = code.InternalServerError
	}

	log.Println("example:", raError)

	response.WithStatus(statusCode)
	response.Write(rsp.ErrorJson(c, message))
}

func getRuntimeStack() string {
	buf := make([]byte, 5120)
	n := runtime.Stack(buf, false)

	return string(buf[:n])
}
