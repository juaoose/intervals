package infra

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreatePool() (*pgxpool.Pool, error) {
	dbUrl, dbUrlErr := getDbUrl()
	if dbUrlErr != nil {
		return nil, dbUrlErr
	}
	dbpool, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		return nil, err
	}
	return dbpool, err
}

func getDbUrl() (string, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl != "" {
		return dbUrl, nil
	}
	return dbUrl, errors.New("db url was not configured")
}
