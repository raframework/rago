package users

import (
	"github.com/raframework/rago/rahttp"
)

type Password struct {
}

func (p *Password) Update(request *rahttp.Request, response *rahttp.Response) {
	data := "example: Password.Updating..."

	response.WithStatus(200).Write(data)
}
