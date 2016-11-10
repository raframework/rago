package resource

import (
	"log"

	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/validation"
)

type Users struct {
}

func (u *Users) Create(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Create...")

	parsedBody := request.GetParsedBody()
	rules := map[string]interface{}{
		"username": "required|email",
		"password": "required",
	}

	log.Println("example: parsedBody: ", parsedBody)

	validator := validation.New(parsedBody, rules)
	if validator.Fails() {
		panic(validator.GetMessage())
	}

	// TODO: do some register steps

	response.WithStatus(201).Write(`{"id": 1, "email": "test@gmail.com"}`)
}

func (u *Users) Update(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Update...")
}

func (u *Users) Get(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Get...")
}

func (u *Users) Delete(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Delete...")
}

func (u *Users) List(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.List...")
}
