package main

import (
	"fmt"
	"os"

	"github.com/pseudo-su/golang-service-template/internal"
	"github.com/pseudo-su/golang-service-template/internal/config"
)

func main() {
	appCtx, err := config.NewApplicationContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to db %s", err.Error())
		os.Exit(1)
	}

	server := internal.Bootstrap(appCtx)

	server.Start(appCtx.ServerPort())
}
