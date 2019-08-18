package config

import (
	"fmt"
)

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

//Version returns application version
func (cfg *ApplicationConfig) Version() string {
	return cfg.envValues.APIVersion
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
	return fmt.Sprintf("/%s/%s", cfg.envValues.ServiceURIName, cfg.envValues.APIVersion)
}
