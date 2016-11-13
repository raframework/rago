package handler

import (
	"fmt"
	"log"
	"reflect"

	"github.com/raframework/rago/example/app/config/code"
	"github.com/raframework/rago/example/app/lib/apperror"
	"github.com/raframework/rago/example/app/lib/rsp"
	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/rahttp"
)

type errorResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func ErrorHandler(err interface{}, request *rahttp.Request, response *rahttp.Response) {

	log.Println("example: err type: ", reflect.ValueOf(err).Type().Kind())
	log.Println("example: errorHandler with err: ", err)

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

	response.WithStatus(500)
	response.Write(rsp.Json(errorResponse{code.InternalServerError, fmt.Sprint("Internal server error with: ", err), ""}))
}

func handleAppError(appError *apperror.AppError, request *rahttp.Request, response *rahttp.Response) {
	statusCode := appError.Typ()
	c := appError.Code()

	var message string
	message = appError.Message()
	if message == "" {
		message = code.Message(c)
	}

	response.WithStatus(statusCode)
	response.Write(rsp.Json(errorResponse{c, message, ""}))
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

	case raerror.TypBadBody:
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

	if message == "" {
		message = code.Message(c)
	}

	response.WithStatus(statusCode)
	response.Write(rsp.Json(errorResponse{c, message, ""}))
}
