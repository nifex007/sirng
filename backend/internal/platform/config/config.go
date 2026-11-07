package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{ Port string }

func Load() (Config, error) {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return Config{}, fmt.Errorf("PORT must be an integer between 1 and 65535")
	}
	return Config{Port: strconv.Itoa(n)}, nil
}
