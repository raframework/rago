package resource

import (
	"github.com/raframework/rago/example/app/lib/rsp"
	"github.com/raframework/rago/rahttp"
)

type Users struct {
}

func (u *Users) Create(request *rahttp.Request, response *rahttp.Response) {
	parsedBody := request.GetParsedBody()

	// TODO: do some register steps

	result := struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}{
		1,
		parsedBody["username"].(string),
	}

	response.WithStatus(201).Write(rsp.Json(result))
}

func (u *Users) Update(request *rahttp.Request, response *rahttp.Response) {
	data := "example: Users.Update..."

	response.WithStatus(200).Write(data)
}

func (u *Users) Get(request *rahttp.Request, response *rahttp.Response) {
	data := "example: Users.Get..."

	response.WithStatus(200).Write(data)
}

func (u *Users) Delete(request *rahttp.Request, response *rahttp.Response) {
	data := "example: Users.Delete..."

	response.WithStatus(204).Write(data)
}

func (u *Users) List(request *rahttp.Request, response *rahttp.Response) {
	data := "example: Users.List..."

	response.WithStatus(200).Write(data)
}
