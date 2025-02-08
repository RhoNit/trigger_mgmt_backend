package common

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDatabase() (*pgx.Conn, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABSE_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s/?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
