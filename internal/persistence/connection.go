package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func CloseDB(db *sql.DB) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func ConnectToDb(ctx context.Context) (*sql.DB, error) {
	connectionCfg, err := newConnectionCfgFromConfig(ctx)
	if err != nil {
		return nil, err
	}
	return connectToDbWithConnectionCfg(ctx, connectionCfg)
}

// Private.
type connectionCfg struct {
	DbHost   string `env:"DB_HOST" envDefault:"localhost"`
	Username string `env:"DB_USER_NAME" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"1234"`
	DbName   string `env:"DB_NAME" envDefault:"golang_service_template_localdev"`
}

func newConnectionCfgFromConfig(ctx context.Context) (*connectionCfg, error) {
	cfg := &connectionCfg{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("cannot find env config for db: %v \n", err)
	}

	switch os.Getenv("DB_CONFIG_MODE") {
	case "env_vars", "":
		return cfg, nil
	default:
		return nil, fmt.Errorf("Invalid value supplied to `DB_CONFIG_MODE` must be `env_vars`")
	}
}

// Private.
func connectToDbWithConnectionCfg(ctx context.Context, in *connectionCfg) (*sql.DB, error) {
	dbURI := fmt.Sprintf(
		"host=%s user=%s sslmode=disable password=%s dbname=%s port=5432",
		in.DbHost,
		in.Username,
		in.Password,
		in.DbName,
	)
	return sql.Open("pgx", dbURI)
}
