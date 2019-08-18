package main

import (
	"github.com/pseudo-su/golang-service-template/internal"
	"github.com/pseudo-su/golang-service-template/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	internal.Bootstrap(cfg)
}
