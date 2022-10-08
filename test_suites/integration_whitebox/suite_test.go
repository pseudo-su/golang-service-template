package integration

import (
	"context"
	"database/sql"
	"testing"

	"github.com/pseudo-su/golang-service-template/internal/persistence"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	sqlDB            *sql.DB
	migrationManager *persistence.MigrationManager
}

//nolint:stylecheck
func (s *TestSuite) SetupSuite() {

	sqlDB, err := persistence.ConnectToDb(context.Background())
	if err != nil {
		panic(err)
	}

	migrationManager, err := persistence.NewMigrationManager(sqlDB)
	if err != nil {
		s.Suite.FailNow("Error: creating migration manager", err)
	}

	s.sqlDB = sqlDB
	s.migrationManager = migrationManager
}

func (s *TestSuite) TeardownSuite() {
	persistence.CloseDB(s.sqlDB)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) BeforeTest(suiteName, testName string) {
	err := s.migrationManager.MigrateUpAll()
	if err != nil {
		s.Suite.FailNow("Error: DB Migration didn't work", err)
	}
}

func (s *TestSuite) AfterTest(suiteName, testName string) {
	err := s.migrationManager.Unsafe_MigrateDownAll()
	if err != nil {
		s.Suite.FailNow("Error: DB Delete All Cleanup after test", err)
	}
}
