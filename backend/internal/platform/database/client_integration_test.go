//go:build integration

package database

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestConnectionIntegration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Fatal("DATABASE_URL is required for integration tests")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := Open(ctx, databaseURL)
	if err != nil {
		t.Fatal("cannot configure Postgres client; verify DATABASE_URL")
	}
	defer pool.Close()
	if err := Check(ctx, pool); err != nil {
		t.Fatal("Postgres SELECT 1 failed; verify availability, credentials, and database")
	}
}
