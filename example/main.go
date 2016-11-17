package main

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config/res"
	"github.com/raframework/rago/example/app/lib/handler"
	"github.com/raframework/rago/example/app/processor"
	"github.com/raframework/rago/ralog"
)

func init() {
	ralog.SetLevel(ralog.LDebug)
}

func main() {
	app := rago.NewApp(res.UriPatterns)

	requestHandler := func(c *rago.Context) {
		c.WithErrorHandler(handler.ErrorHandler)
		c.Call(&processor.Predo{})
		c.MatchUriPattern()
		c.CallResourceAction()
		c.Respond()
	}

	app.WithRequestHanlder(requestHandler).Run(":8800")
}
