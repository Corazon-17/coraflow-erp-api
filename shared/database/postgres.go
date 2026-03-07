package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
