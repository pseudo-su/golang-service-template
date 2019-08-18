package smoke

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/pseudo-su/golang-service-template/pkg"
	"github.com/sirupsen/logrus"
)

func buildBaseURL(cfg *TestSuiteConfig) string {
	path := fmt.Sprintf("/%s/%s", cfg.ServiceURIName, cfg.APIVersion)
	return fmt.Sprintf("%s://%s:%s%s", cfg.APIScheme, cfg.APIHost, cfg.APIPort, path)
}

func smokeSuiteSetup() (*TestSuiteConfig, *pkg.ClientWithResponses) {
	cfg := ParseSuiteConfig()
	serverBaseURL := buildBaseURL(cfg)
	apiClient := pkg.NewClientWithResponses(serverBaseURL)
	return cfg, apiClient
}

type TestSuiteConfig struct {
	APIScheme      string `env:"API_SCHEME" envDefault:"http"`
	APIHost        string `env:"API_HOST" envDefault:"localhost"`
	APIPort        string `env:"API_PORT" envDefault:"3000"`
	ServiceURIName string `env:"SERVICE_URI_NAME" envDefault:"golang-service-template"`
	APIVersion     string `env:"API_VERSION" envDefault:"v1"`
	BasePath       string `env:"BASE_PATH"`
	ServiceToken   string `env:"SERVICE_TOKEN"`
	Env            string `env:"ENV" envDefault:"local"`
	LogLevel       string `env:"LOG_LEVEL"`
}

func ParseSuiteConfig() *TestSuiteConfig {
	cfg := &TestSuiteConfig{}
	if err := env.Parse(cfg); err != nil {
		logrus.Fatalf("unable to find env var key: %v \n", err)
	}
	return cfg
}
