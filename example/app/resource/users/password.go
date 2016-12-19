package users

import (
	"github.com/coderd/glog"
	"github.com/raframework/rago/rahttp"
)

type Password struct {
}

func (p *Password) Update(request *rahttp.Request, response *rahttp.Response) {
	glog.Debug("example: Password.List...")
}
