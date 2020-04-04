package smoke

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env"
	"github.com/pseudo-su/golang-service-template/internal"
	"github.com/pseudo-su/golang-service-template/internal/config"
	"github.com/pseudo-su/golang-service-template/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

//nolint:go-lint
type TestSuite struct {
	suite.Suite
	server    *config.Server
	suiteCtx  *TestSuiteContext
	apiClient *pkg.Client
}

func (suite *TestSuite) SetupSuite() {
	suite.suiteCtx = ParseSuiteConfig()
	serverBaseURL := buildBaseURL(suite.suiteCtx)
	apiClient, err := pkg.NewClient(serverBaseURL)
	if err != nil {
		panic(err)
	}
	suite.apiClient = apiClient
	suite.server = internal.Bootstrap(suite.suiteCtx)

	if suite.suiteCtx.envValues.UseEmbeddedServer {
		go func() {
			err := suite.server.ListenAndServe()
			if err != nil {
				log.Infof("Listen and serve: %v", err)
			}
		}()
	}
}

func (suite *TestSuite) TeardownSuite() {
	suite.server.ShutdownReq <- true
}

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
	UseEmbeddedServer bool   `env:"USE_EMBEDDED_SERVER" envDefault:"true"`
	APIScheme         string `env:"API_SCHEME" envDefault:"http"`
	APIHost           string `env:"API_HOST" envDefault:"localhost"`
	APIPort           int    `env:"API_PORT" envDefault:"3000"`
	APIBasepath       string `env:"API_BASEPATH" envDefault:"/golang-service-template/v1"`
	Env               string `env:"ENV" envDefault:"local"`
	LogLevel          string `env:"LOG_LEVEL" envDefault:"debug"`
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
