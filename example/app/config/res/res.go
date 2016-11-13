package res

import (
	"github.com/raframework/rago/example/app/resource"
	"github.com/raframework/rago/example/app/resource/users"
	"github.com/raframework/rago/rahttp"
)

var UriPatterns = map[rahttp.UriPattern]rahttp.ResourceAndMethod{
	"/users": {
		resource.Users{},
		[]rahttp.Method{"POST", "GET"},
	},
	"/users/:id": {
		resource.Users{},
		[]rahttp.Method{"GET", "PUT", "DELETE"},
	},
	"/users/:id/password": {
		users.Password{},
		[]rahttp.Method{"PUT"},
	},
}
