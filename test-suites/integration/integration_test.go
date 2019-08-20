package integration

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env"
	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

//nolint
type TestSuite struct {
	suite.Suite
	server    *config.Server
	cfg       *TestSuiteConfig
	apiClient *pkg.ClientWithResponses
}

func (suite *TestSuite) SetupSuite() {
	suite.cfg = ParseSuiteConfig()
	serverBaseURL := buildBaseURL(suite.cfg)
	suite.apiClient = pkg.NewClientWithResponses(serverBaseURL)

	// suite.server = internal.Bootstrap(suite.cfg)
	// if suite.cfg.envValues.UseEmbeddedServer {
	// 	go func() {
	// 		err := suite.server.ListenAndServe()
	// 		if err != nil {
	// 			log.Printf("Listen and serve: %v", err)
	// 		}
	// 	}()
	// 	time.Sleep(2 * time.Second)
	// }
}

func (suite *TestSuite) TeardownSuite() {
	// suite.server.ShutdownReq <- true
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func buildBaseURL(cfg *TestSuiteConfig) string {
	basepath := cfg.ServiceBasepath()
	return fmt.Sprintf(
		"%s://%s:%v%s",
		cfg.envValues.APIScheme,
		cfg.envValues.APIHost,
		cfg.envValues.APIPort,
		basepath,
	)
}

type TestSuiteConfig struct {
	envValues *testSuiteEnv
}

func (cfg *TestSuiteConfig) Env() string {
	return cfg.envValues.Env
}

func (cfg *TestSuiteConfig) ServerPort() int {
	return cfg.envValues.APIPort
}

func (cfg *TestSuiteConfig) ServiceBasepath() string {
	return fmt.Sprintf("/%s/%s", cfg.envValues.ServiceURIName, cfg.envValues.APIVersion)
}

type testSuiteEnv struct {
	UseEmbeddedServer bool   `env:"USE_EMBEDDED_SERVER" envDefault:"true"`
	APIScheme         string `env:"API_SCHEME" envDefault:"http"`
	APIHost           string `env:"API_HOST" envDefault:"localhost"`
	APIPort           int    `env:"API_PORT" envDefault:"3000"`
	ServiceURIName    string `env:"SERVICE_URI_NAME" envDefault:"golang-service-template"`
	APIVersion        string `env:"API_VERSION" envDefault:"v1"`
	Env               string `env:"ENV" envDefault:"local"`
	LogLevel          string `env:"LOG_LEVEL" envDefault:"debug"`
}

func ParseSuiteConfig() *TestSuiteConfig {
	envValues := &testSuiteEnv{}
	if err := env.Parse(envValues); err != nil {
		log.Fatalf("unable to find env var key: %v \n", err)
	}
	return &TestSuiteConfig{
		envValues: envValues,
	}
}
