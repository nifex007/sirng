//go:build integration

package objectstore

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/nifex007/sirng/backend/internal/platform/config"
)

func TestConnectionIntegration(t *testing.T) {
	cfg := config.Connections{Region: os.Getenv("AWS_REGION"), Bucket: os.Getenv("S3_BUCKET"), Endpoint: os.Getenv("S3_ENDPOINT")}
	if cfg.Region == "" || cfg.Bucket == "" {
		t.Fatal("AWS_REGION and S3_BUCKET are required for integration tests")
	}
	if value, ok := os.LookupEnv("S3_USE_PATH_STYLE"); ok {
		var err error
		cfg.PathStyle, err = strconv.ParseBool(value)
		if err != nil {
			t.Fatal("S3_USE_PATH_STYLE must be a boolean")
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := New(ctx, cfg)
	if err != nil {
		t.Fatal("cannot configure S3 client; verify region and AWS credentials")
	}
	if err := Check(ctx, client, cfg.Bucket); err != nil {
		t.Fatal("S3 HeadBucket failed; verify endpoint, credentials, bucket, and permissions")
	}
}
