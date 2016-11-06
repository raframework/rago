package users

import (
	"log"

	"github.com/raframework/rago/rahttp"
)

type Password struct {
}

func (p *Password) Update(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Password.List...")
}
