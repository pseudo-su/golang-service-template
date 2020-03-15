package main

import (
	"github.com/pseudo-su/golang-service-template/internal"
	"github.com/pseudo-su/golang-service-template/internal/config"
)

func main() {
	appCtx := config.NewApplicationContext()
	server := internal.Bootstrap(appCtx)

	server.Start(appCtx.ServerPort())
}
