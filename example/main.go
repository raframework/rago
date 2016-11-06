package main

import (
	"log"

	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config"
)

func errorHanlder(err interface{}) {
	log.Println("example: errorHandler with ", err)
	panic(err)
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
