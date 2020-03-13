package config

type ApplicationConfig struct {
	envValues *envConfig
}

//NewApplicationConfig loads config values from environment and initializes config
func NewApplicationConfig() *ApplicationConfig {
	envValues := newEnvironmentConfig()

	return &ApplicationConfig{
		envValues: envValues,
	}
}

//ServerPort returns the port no to listen for requests
func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}

//ServiceBasepath returns the authorisation scope prefix
func (cfg *ApplicationConfig) Env() string {
	return cfg.envValues.Env
}

//ServiceBasepath returns the authorisation scope prefix
func (cfg *ApplicationConfig) ServiceBasepath() string {
	return cfg.envValues.APIBasepath
}
