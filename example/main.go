package main

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config"
)

func main() {
	app := rago.NewApp(config.UriPatterns)

	requestHandler := func(c *rago.Context) {
		c.MatchUriPattern()
		c.Call()
		c.CallResourceAction()
		c.Respond()
	}

	app.WithRequestHanlder(requestHandler).Run(":8800")
}
