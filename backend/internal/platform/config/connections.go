package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Connections struct {
	DatabaseURL, Region, Bucket, Endpoint string
	PathStyle                             bool
}

func LoadConnections() (Connections, error) {
	c := Connections{DatabaseURL: os.Getenv("DATABASE_URL"), Region: os.Getenv("AWS_REGION"), Bucket: os.Getenv("S3_BUCKET"), Endpoint: os.Getenv("S3_ENDPOINT")}
	for _, name := range []string{"DATABASE_URL", "AWS_REGION", "S3_BUCKET"} {
		if os.Getenv(name) == "" {
			return c, fmt.Errorf("%s is required", name)
		}
	}
	u, err := url.Parse(c.DatabaseURL)
	if err != nil || u.Host == "" || (u.Scheme != "postgres" && u.Scheme != "postgresql") {
		return c, fmt.Errorf("DATABASE_URL must be a Postgres URL")
	}
	if c.Endpoint != "" {
		u, err := url.Parse(c.Endpoint)
		if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") || u.User != nil {
			return c, fmt.Errorf("S3_ENDPOINT must be an HTTP or HTTPS URL without credentials")
		}
	}
	if value, ok := os.LookupEnv("S3_USE_PATH_STYLE"); ok {
		c.PathStyle, err = strconv.ParseBool(value)
		if err != nil {
			return c, fmt.Errorf("S3_USE_PATH_STYLE must be a boolean")
		}
	}
	return c, nil
}
