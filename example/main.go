package main

import (
	"log"

	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config"
	"github.com/raframework/rago/rahttp"
)

func errorHanlder(err interface{}, request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: errorHandler with ", err)
	response.WithStatus(500)
	response.Write(`{"code": 10001, "message": "Internel Server Error"}`)
}

func main() {
	app := rago.NewApp(config.UriPatterns)

	requestHandler := func(c *rago.Context) {
		c.WithErrorHandler(errorHanlder)
		c.MatchUriPattern()
		c.CallResourceAction()
		c.Respond()
	}

	app.WithRequestHanlder(requestHandler).Run(":8800")
}
