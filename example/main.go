package main

import (
	"github.com/raframework/rago"
	"github.com/raframework/rago/example/app/config"
)

func main() {
	app := rago.NewApp(config.UriPatterns)

	app.MatchUriPattern().CallResourceAction().Respond()
}
