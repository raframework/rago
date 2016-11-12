package rago

import (
	"github.com/raframework/rago/rahttp"
)

type Processor interface {
	Process(*rahttp.Request, *rahttp.Response)
}
