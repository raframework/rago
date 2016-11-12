package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/raframework/rago/example/app/config/code"
	"github.com/raframework/rago/example/app/lib/apperror"
	"github.com/raframework/rago/rahttp"
)

type errorResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func ErrorHandler(err interface{}, request *rahttp.Request, response *rahttp.Response) {

	log.Println("err type: ", reflect.ValueOf(err).Type().Kind())
	log.Println("example: errorHandler with ", err)

	appError, ok := err.(*apperror.AppError)
	if ok {
		statusCode := appError.Typ()
		c := appError.Code()
		var message string
		message = appError.Message()
		if message == "" {
			message = code.Message(c)
		}

		response.WithStatus(statusCode)
		rsp, _ := json.Marshal(errorResponse{c, message, ""})
		response.Write(string(rsp))
		return
	}

	response.WithStatus(500)
	rsp, _ := json.Marshal(errorResponse{code.InternalServerError, fmt.Sprint("Internal server error with: ", err), ""})
	response.Write(string(rsp))
}
