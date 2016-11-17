package users

import (
	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/ralog"
)

type Password struct {
}

func (p *Password) Update(request *rahttp.Request, response *rahttp.Response) {
	ralog.Debug("example: Password.List...")
}
