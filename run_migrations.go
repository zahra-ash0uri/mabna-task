package main

import (
	"fmt"

	"github.com/jackc/pgx/v4/migrate"
	"github.com/jackc/pgx/v4/migrate/source/file"
)

// conn, err := models.ConnectToPostgres()
//
//	if err != nil {
//	    fmt.Println("Connection to DB failed...")
//	}
//
// defer conn.Close(context.Background())
func RunMigrations() {

	sourceDriver, err := file.WithInstance("file://migrations")
	if err != nil {
		fmt.Println("Can not get source driver!")
	}
	m, err := migrate.NewWithInstance("file", sourceDriver, "database_name_here", &migrate.Options{})
	if err != nil {
		fmt.Println("Can not create migration instance!")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	fmt.Println("Migrations applied successfully")

}
