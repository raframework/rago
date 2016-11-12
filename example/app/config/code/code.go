package code

import (
	"fmt"

	"github.com/raframework/rago/example/app/config"
)

const (
	ParamError           = 10001
	ResourceNotFound     = 10002
	BadBody              = 10003
	MethodNotAllowed     = 10004
	UnsupportedMediaType = 10005
	InternalServerError  = 10006
	PermissionDenied     = 10007
)

var messages = map[int]map[string]string{
	ParamError: {
		config.LangEn: "Parameters error",
	},
	ResourceNotFound: {
		config.LangEn: "Resource not found",
	},
	BadBody: {
		config.LangEn: "Bad HTTP body",
	},
	MethodNotAllowed: {
		config.LangEn: "HTTP method not allowed",
	},
	UnsupportedMediaType: {
		config.LangEn: "Unsupported media type",
	},
	InternalServerError: {
		config.LangEn: "Internal server error",
	},
	PermissionDenied: {
		config.LangEn: "Permisson denied",
	},
}

func Message(code int) string {
	lang := config.Lang()
	message, ok := messages[code][lang]
	if !ok {
		panic(fmt.Sprintf("config: code %d not exists", code))
	}

	return message
}
