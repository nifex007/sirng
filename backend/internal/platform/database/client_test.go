package database

import (
	"context"
	"testing"
)

func TestOpenRejectsInvalidConfiguration(t *testing.T) {
	pool, err := Open(context.Background(), "postgres://localhost/db?pool_max_conns=invalid")
	if pool != nil {
		pool.Close()
	}
	if err == nil {
		t.Fatal("expected invalid pool configuration to fail")
	}
}
