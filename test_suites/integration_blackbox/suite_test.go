package integration

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env"
	"github.com/pseudo-su/golang-service-template/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	suiteCtx  *TestSuiteContext
	apiClient *pkg.ClientWithResponses
}

//nolint:stylecheck
func (s *TestSuite) SetupSuite() {
	s.suiteCtx = ParseSuiteConfig()
	serverBaseURL := buildBaseURL(s.suiteCtx)

	apiClient, err := pkg.NewClientWithResponses(serverBaseURL)
	if err != nil {
		panic(err)
	}

	s.apiClient = apiClient
}

func (s *TestSuite) TeardownSuite() {}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func buildBaseURL(suiteCtx *TestSuiteContext) string {
	basepath := suiteCtx.ServiceBasepath()

	return fmt.Sprintf(
		"%s://%s:%v%s",
		suiteCtx.envValues.APIScheme,
		suiteCtx.envValues.APIHost,
		suiteCtx.envValues.APIPort,
		basepath,
	)
}

type TestSuiteContext struct {
	envValues *testSuiteEnv
}

func (suiteCtx *TestSuiteContext) Env() string {
	return suiteCtx.envValues.Env
}

func (suiteCtx *TestSuiteContext) ServerPort() int {
	return suiteCtx.envValues.APIPort
}

func (suiteCtx *TestSuiteContext) ServiceBasepath() string {
	return suiteCtx.envValues.APIBasepath
}

type testSuiteEnv struct {
	APIScheme   string `env:"API_SCHEME" envDefault:"http"`
	APIHost     string `env:"API_HOST" envDefault:"localhost"`
	APIPort     int    `env:"API_PORT" envDefault:"3000"`
	APIBasepath string `env:"API_BASEPATH" envDefault:"/golang-service-template/v1"`
	Env         string `env:"ENV" envDefault:"local"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
}

func ParseSuiteConfig() *TestSuiteContext {
	envValues := &testSuiteEnv{}
	if err := env.Parse(envValues); err != nil {
		log.Fatalf("unable to find env var key: %v \n", err)
	}

	return &TestSuiteContext{
		envValues: envValues,
	}
}
