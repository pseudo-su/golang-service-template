package persistence

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

type MigrationManager struct {
	db *sql.DB
	m  *migrate.Migrate
}

func NewMigrationManager(db *sql.DB) (*MigrationManager, error) {

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, err
	}

	dbInstance, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithInstance("iofs", d, "mantel_connect_backend", dbInstance)

	if err != nil {
		return nil, err
	}

	return &MigrationManager{
		db: db,
		m:  m,
	}, nil
}

func (mg *MigrationManager) ForceMigrationVersion(version *int64) error {
	return fmt.Errorf(
		"MigrationManager.ForceMigrationVersion(version: %v) not implemented", version,
	)
}

func (mg *MigrationManager) MigrateToVersion(version *int64) error {
	return fmt.Errorf(
		"MigrationManager.MigrateToVersion(version: %v) not implemented", version,
	)
}

func (mg *MigrationManager) MigrateUpAll() error {
	err := mg.m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}
	return nil
}

func (mg *MigrationManager) Unsafe_MigrateDownAll() error {
	err := mg.m.Down()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}
	return nil
}

func (mg *MigrationManager) MigrateUpBy(count *int) error {
	return fmt.Errorf(
		"MigrationManager.MigrateUpBy(count: %v) not implemented", count,
	)
}

func (mg *MigrationManager) MigrateDownBy(count *int) error {
	return fmt.Errorf(
		"MigrationManager.MigrateDownBy(count: %v) not implemented", count,
	)
}
