package resource

import (
	"log"
	"strconv"

	"github.com/raframework/rago/example/app/config/code"
	"github.com/raframework/rago/example/app/lib/apperror"
	"github.com/raframework/rago/example/app/lib/rsp"
	"github.com/raframework/rago/rahttp"
	"github.com/raframework/rago/validation"
)

type Users struct {
}

func (u *Users) Create(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Create...")

	queryParams := request.GetQueryParams()
	parsedBody := request.GetParsedBody()
	rules := map[string]interface{}{
		"username": "required|string|email",
		"password": "required|string",
	}

	log.Println("example: queryParams: ", queryParams)
	log.Println("example: parsedBody: ", parsedBody)

	validator := validation.New(parsedBody, rules)
	if validator.Fails() {
		apperror.PanicWithMessage(apperror.BadRequest, code.ParamError, validator.GetMessage())
	}

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
	log.Println("example: Users.Update...")

	attributes := request.GetAttributes()
	queryParams := request.GetQueryParams()
	parsedBody := request.GetParsedBody()
	rules := map[string]interface{}{
		"username": "required|string|email",
		"password": "string",
		"age":      "float|min:1|max:100",
		"score":    "float|between: 1, 100",
		"size":     "size: 88",
	}

	log.Println("example: attributes: ", attributes)
	log.Println("example: queryParams: ", queryParams)
	log.Println("example: parsedBody: ", parsedBody)

	id, err := strconv.Atoi(attributes["id"])
	if err != nil {
		apperror.PanicWithMessage(apperror.BadRequest, code.ParamError, "Bad URL")
	}

	validator := validation.New(parsedBody, rules)
	if validator.Fails() {
		apperror.PanicWithMessage(apperror.BadRequest, code.ParamError, validator.GetMessage())
	}

	// TODO: do some updating steps

	result := struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}{
		id,
		parsedBody["username"].(string),
	}

	response.WithStatus(200).Write(rsp.Json(result))
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
