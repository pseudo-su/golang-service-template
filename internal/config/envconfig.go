package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type envConfig struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"golang-service-template"`
	APIBasepath string `env:"API_BASEPATH" envDefault:"/golang-service-template/v1"`
	ServerPort  int    `env:"SERVER_PORT" envDefault:"3000"`
	Env         string `env:"ENV" envDefault:"local"`

	// Logs
	LogLevel  string `env:"LOG_LEVEL"`
	LogFormat string `env:"LOG_FORMAT"`
}

func newEnvironmentConfig() *envConfig {
	cfg := &envConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("cannot find configs for server: %v \n", err)
	}
	return cfg
}
