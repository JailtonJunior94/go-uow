package migrate

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	postgresmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jailtonjunior94/go-uow/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrMigrateVersion        = errors.New("error checking migration version")
	ErrPostgresMigrateDriver = errors.New("unable to instantiate postgres migration driver")
)

type Migrate struct {
	migrate *migrate.Migrate
	logger  logger.Logger
}

func NewMigrate(logger logger.Logger, db *sql.DB, migratePath, dbName string) (*Migrate, error) {
	if db == nil {
		return nil, ErrPostgresMigrateDriver
	}

	driver, err := postgresmigrate.WithInstance(db, &postgresmigrate.Config{})
	if err != nil {
		return nil, ErrPostgresMigrateDriver
	}

	m, err := migrate.NewWithDatabaseInstance(migratePath, dbName, driver)
	if err != nil {
		return nil, err
	}
	return &Migrate{migrate: m, logger: logger}, nil
}

func (m *Migrate) ExecuteMigration() error {
	version, isDirty, err := m.migrate.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return ErrMigrateVersion
	}
	m.logger.Infof("running migration scripts... migration_version: %d is_dirty: %t", version, isDirty)

	err = m.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		m.logger.Infof("no database updates")
		return nil
	}

	if err != nil {
		return err
	}
	return nil
}
