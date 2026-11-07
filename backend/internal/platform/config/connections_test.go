package config

import "testing"

func TestConnections(t *testing.T) {
	for k, v := range map[string]string{"DATABASE_URL": "postgres://user:password@localhost/db", "AWS_REGION": "us-east-1", "S3_BUCKET": "test", "S3_ENDPOINT": "http://localhost:4566", "S3_USE_PATH_STYLE": "true"} {
		t.Setenv(k, v)
	}
	cfg, err := LoadConnections()
	if err != nil || !cfg.PathStyle {
		t.Fatalf("%v", err)
	}
	for k, v := range map[string]string{"DATABASE_URL": "secret", "AWS_REGION": "", "S3_BUCKET": "", "S3_ENDPOINT": "bad", "S3_USE_PATH_STYLE": "bad"} {
		t.Run(k, func(t *testing.T) {
			t.Setenv(k, v)
			if _, err := LoadConnections(); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}
