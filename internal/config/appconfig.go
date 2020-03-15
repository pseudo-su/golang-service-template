package config

type ApplicationContext struct {
	envValues *envConfig
}

//NewApplicationContext loads config values from environment and initializes config
func NewApplicationContext() *ApplicationContext {
	envValues := newEnvironmentConfig()

	return &ApplicationContext{
		envValues: envValues,
	}
}

//ServerPort returns the port no to listen for requests
func (appCtx *ApplicationContext) ServerPort() int {
	return appCtx.envValues.ServerPort
}

//ServiceBasepath returns the authorisation scope prefix
func (appCtx *ApplicationContext) Env() string {
	return appCtx.envValues.Env
}

//ServiceBasepath returns the authorisation scope prefix
func (appCtx *ApplicationContext) ServiceBasepath() string {
	return appCtx.envValues.APIBasepath
}
