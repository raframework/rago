package config

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/resource"
	"github.com/raframework/rago/example/app/resource/users"
)

var UriPatterns = rago.UriPatterns{
	"/users": {
		&resource.Users{},
		{"POST", "GET", "PUT", "DELETE"},
	},
	"/users/:id/password": {
		&users.Password{},
		{"PUT"},
	},
}
