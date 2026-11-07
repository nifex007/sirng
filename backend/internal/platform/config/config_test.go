package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Setenv("PORT", "")
	os.Unsetenv("PORT")
	cfg, err := Load()
	if err != nil || cfg.Port != "8080" {
		t.Fatalf("default: %v %v", cfg, err)
	}
	for _, port := range []string{"1", "8081", "65535", "", "abc", "0", "65536", "-1"} {
		t.Run(port, func(t *testing.T) {
			t.Setenv("PORT", port)
			_, err := Load()
			valid := port == "1" || port == "8081" || port == "65535"
			if (err == nil) != valid {
				t.Fatalf("PORT=%q: %v", port, err)
			}
		})
	}
}
