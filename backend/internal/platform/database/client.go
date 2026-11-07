package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Open(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, databaseURL)
}

func Check(ctx context.Context, pool *pgxpool.Pool) error {
	var value int
	if err := pool.QueryRow(ctx, "SELECT 1").Scan(&value); err != nil {
		return err
	}
	if value != 1 {
		return fmt.Errorf("unexpected database probe result")
	}
	return nil
}
