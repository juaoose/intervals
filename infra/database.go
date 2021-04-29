package infra

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreatePool() (*pgxpool.Pool, error) {
	//TODO fix database URL
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}
	return dbpool, err
}
