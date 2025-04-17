package config

import "testing"

func TestConfigHasValues(t *testing.T) {
	cfg := Load()

	if cfg.ServerAddress == "" {
		t.Error("Server address should not be empty")
	}

	if cfg.DatabaseURL == "" {
		t.Error("Database URL should not be empty")
	}
}
