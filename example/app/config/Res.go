package config

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/resource"
	"github.com/raframework/rago/example/app/resource/users"
)

var UriPatterns = map[rago.UriPattern]rago.ResourceMethod{
	"/users": {
		Resource: &resource.Users{},
		Methods:  {"POST", "GET", "PUT", "DELETE"},
	},
	"/users/:id/password": {
		Resource: &users.Password{},
		Methods:  {"PUT"},
	},
}
