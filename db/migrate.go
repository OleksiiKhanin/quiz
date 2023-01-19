package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateSchema(storage *sql.DB, migrateFilesPath string) error {
	driver, err := postgres.WithInstance(storage, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("get driver for migration: %w", err)
	}
	path := fmt.Sprintf("file://%s", migrateFilesPath)
	log.Printf("Try find migrate files in %s\n", path)
	m, err := migrate.NewWithDatabaseInstance(
		path,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up failed: %w", err)
	}
	return nil
}
