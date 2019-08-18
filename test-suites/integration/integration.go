package integration

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/sirupsen/logrus"
)

func smokeSuiteSetup() (*TestSuiteConfig, *pkg.ClientWithResponses) {
	cfg := ParseSuiteConfig()
	serverBaseURL := fmt.Sprintf(
		"%v/%v/%v",
		cfg.Host,
		cfg.ServiceVersion,
		cfg.BasePath,
	)
	apiClient := pkg.NewClientWithResponses(serverBaseURL)
	return cfg, apiClient
}

type TestSuiteConfig struct {
	Host           string `env:"HOST"`
	ServiceVersion string `env:"SERVICE_VERSION"`
	BasePath       string `env:"BASE_PATH"`
	ServiceToken   string `env:"SERVICE_TOKEN"`
}

func ParseSuiteConfig() *TestSuiteConfig {
	cfg := &TestSuiteConfig{}
	if err := env.Parse(cfg); err != nil {
		logrus.Fatalf("unable to find env var key: %v \n", err)
	}
	return cfg
}
