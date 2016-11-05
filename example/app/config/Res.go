package config

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/resource"
	"github.com/raframework/rago/example/app/resource/users"
)

var UriPatterns = map[rago.UriPattern]rago.ResourceMethod{
	"/users": {
		&resource.Users{},
		[]rago.Method{"POST", "GET", "PUT", "DELETE"},
	},
	"/users/:id/password": {
		&users.Password{},
		[]rago.Method{"PUT"},
	},
}
