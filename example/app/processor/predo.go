package processor

import (
	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/ralog"
)

type Predo struct {
}

func (p *Predo) Process(request *rahttp.Request, response *rahttp.Response) {
	ralog.Debug("example.app.processor: predo processing")

	response.WithHeader("Content-Type", "application/json;charset=utf-8")
}
