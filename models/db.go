package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type pgConn struct {
	Conn *pgxpool.Pool
}

var db *pgConn

func New() *pgConn {
	if db == nil {
		dbcon, err := connectToPostgres()
		if err != nil {
			panic(err)
		}
		db = &pgConn{Conn: dbcon}
	}
	return db
}

func connectToPostgres() (*pgxpool.Pool, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting caller information")
		return nil, nil
	}
	rootDir := filepath.Dir(filepath.Dir(filename))
	configPath := filepath.Join(rootDir, ".env")
	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, nil
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)

	Connection, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return Connection, nil

}
