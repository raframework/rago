# Rago

Rago is a RESTful API framework for the Go language. 

## Get Started

1. Download and install it:

    ```bash
    $ go get github.com/raframework/rago
    ```

2. Import it in your code:

    ```go
    import (
    	"github.com/raframework/rago"
    	"github.com/raframework/rago/rahttp"
    )

    ```

3. Example:

    ```bash
    $ cat test.go
    ```

    ```go
    package main

	import (
		"github.com/raframework/rago"
		"github.com/raframework/rago/rahttp"
	)

	// Define the resource.
	type Users struct {
	}

	// Define the resource's action.
	func (u *Users) List(request *rahttp.Request, response *rahttp.Response) {
		data := "List users..."

		response.WithStatus(200).Write(data)
	}

	func main() {
		// Define the routes.
		var uriPatterns = map[rahttp.UriPattern]rahttp.ResourceAndMethod{
			"/users": {
				Users{},
				[]rahttp.Method{"GET"},
			},
		}

		// Create a rago app with the routes given.
		app := rago.NewApp(uriPatterns)

		// Request handler handles the incoming request.
		requestHandler := func(c *rago.Context) {
			c.MatchUriPattern()
			c.CallResourceAction()
			c.Respond()
		}

		// Serves on :8800
		app.WithRequestHanlder(requestHandler).Run(":8800")
	}
    ```

    ```bash
    $ go run test.go
    ```

    Open an another terminal:
    ```bash
    $ curl -v http://127.0.0.1:8800/users
    *   Trying 127.0.0.1...
	* Connected to 127.0.0.1 (127.0.0.1) port 8800 (#0)
	> GET /users HTTP/1.1
	> Host: 127.0.0.1:8800
	> User-Agent: curl/7.49.1
	> Accept: */*
	> 
	< HTTP/1.1 200 OK
	< Date: Sat, 31 Dec 2016 03:00:38 GMT
	< Content-Length: 13
	< Content-Type: text/plain; charset=utf-8
	< 
	* Connection #0 to host 127.0.0.1 left intact
	List users...
    ```
