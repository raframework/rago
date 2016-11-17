package rsp

import (
	"encoding/json"

	"github.com/raframework/rago/example/app/config/code"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Json(v interface{}) string {
	result, _ := json.Marshal(v)

	return string(result)
}

func ErrorJson(c int, message string) string {
	m := code.Message(c)
	if message == "" {
		message = m
	}

	return Json(errorResponse{c, message})
}
