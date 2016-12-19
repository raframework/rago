package resource

import (
	"strconv"

	"github.com/coderd/glog"
	"github.com/raframework/rago/example/app/config/code"
	"github.com/raframework/rago/example/app/lib/apperror"
	"github.com/raframework/rago/example/app/lib/rsp"
	"github.com/raframework/rago/rahttp"
	"github.com/ragopkg/validation"
)

type Users struct {
}

func (u *Users) Create(request *rahttp.Request, response *rahttp.Response) {
	glog.Debug("example: Users.Create...")

	queryParams := request.GetQueryParams()
	parsedBody := request.GetParsedBody()
	rules := map[string]interface{}{
		"username": "required|string|email",
		"password": "required|string",
	}

	glog.Debug("example: queryParams: ", queryParams)
	glog.Debug("example: parsedBody: ", parsedBody)

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
	glog.Debug("example: Users.Update...")

	attributes := request.GetAttributes()
	queryParams := request.GetQueryParams()
	parsedBody := request.GetParsedBody()
	rules := map[string]interface{}{
		"username": "string|email",
		"password": "string",
		"age":      "float|min:1|max:100",
		"score":    "float|between: 1, 100",
		"size":     "size: 88",
	}

	glog.Debug("example: attributes: ", attributes)
	glog.Debug("example: queryParams: ", queryParams)
	glog.Debug("example: parsedBody: ", parsedBody)

	glog.Debug("example: header: X-Test: ", request.GetHeader("X-Test"))

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
	glog.Debug("example: Users.Get...")
}

func (u *Users) Delete(request *rahttp.Request, response *rahttp.Response) {
	glog.Debug("example: Users.Delete...")
}

func (u *Users) List(request *rahttp.Request, response *rahttp.Response) {
	glog.Debug("example: Users.List...")
}
