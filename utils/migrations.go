package utils

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplyMigrations() {
	migrationSource := "file://migrations"
	databaseURL := "postgres://postgres:password@localhost:5432/testgo?sslmode=disable"

	m, err := migrate.New(migrationSource, databaseURL)
	if err != nil {
		log.Fatalf("Error initializing migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
