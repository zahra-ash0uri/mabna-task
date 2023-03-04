package main

import (
	"embed"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/zahra-ash0uri/mabna-task/models"
)

var fs embed.FS

const version = 1

func RunMigrations() {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	db, _ := models.ConnectToPostgres()
	driverInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, "postgresql", driverInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}
