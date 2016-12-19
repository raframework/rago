package main

import (
	"github.com/coderd/glog"
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config/res"
	"github.com/raframework/rago/example/app/lib/handler"
	"github.com/raframework/rago/example/app/processor"
)

func init() {
	glog.SetLevel(glog.LDebug)
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
