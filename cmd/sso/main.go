package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	// Todo: init object config
	config := config.MustLoad()

	fmt.Print(config)

	// Todo: init logger

	// Todo: initialize (app)

	// Todo: run grpc server
}
