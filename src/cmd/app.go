package main

import (
	"fmt"

	"github.com/ayocodingit/boilerplate-clean-archicture-go/src/config"
)

func main() {
	config := config.NewConfig()

	fmt.Println(config)
}
