package main

import (
	"github.com/ayocodingit/boilerplate-clean-archicture-go/src/config"
	"github.com/ayocodingit/boilerplate-clean-archicture-go/src/transport/http"
)

func main() {
	config := config.NewConfig()

	app := http.NewHttp()

	app.Run(config.App.Port.Http)

}
