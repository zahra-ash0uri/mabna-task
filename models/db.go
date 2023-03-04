package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func ConnectToPostgres() (*pgx.Conn, error) {
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

	conn, err := pgx.Connect(context.Background(), connectionString)
	// if conn == nil:
	// 	return nil, err

	return conn, err

}
