package rago

import (
	"github.com/raframework/rago/rahttp"
)

type ErrorHandler func(interface{}, *rahttp.Request, *rahttp.Response)
