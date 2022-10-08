package config

import (
	"context"
	"database/sql"

	"github.com/pseudo-su/golang-service-template/internal/persistence"
)

type ApplicationContext struct {
	ctx       context.Context
	envValues *envConfig
	sqlDB     *sql.DB
}

// NewApplicationContext loads config values from environment and initializes config
func NewApplicationContext() (*ApplicationContext, error) {
	ctx := context.Background()
	envValues := newEnvironmentConfig()

	sqlDB, err := persistence.ConnectToDb(context.Background())
	if err != nil {
		return nil, err
	}

	return &ApplicationContext{
		ctx:       ctx,
		envValues: envValues,
		sqlDB:     sqlDB,
	}, nil
}

// ServerPort returns the port no to listen for requests
func (appCtx *ApplicationContext) ServerPort() int {
	return appCtx.envValues.ServerPort
}

// ServiceBasepath returns the authorisation scope prefix
func (appCtx *ApplicationContext) Env() string {
	return appCtx.envValues.Env
}

// ServiceBasepath returns the authorisation scope prefix
func (appCtx *ApplicationContext) ServiceBasepath() string {
	return appCtx.envValues.APIBasepath
}

// ServiceBasepath returns the authorisation scope prefix
func (appCtx *ApplicationContext) SqlDB() *sql.DB {
	return appCtx.sqlDB
}
