package processor

import (
	"github.com/coderd/glog"
	"github.com/raframework/rago/rahttp"
)

type Predo struct {
}

func (p *Predo) Process(request *rahttp.Request, response *rahttp.Response) {
	glog.Debug("example.app.processor: predo processing")

	response.WithHeader("Content-Type", "application/json;charset=utf-8")
}
