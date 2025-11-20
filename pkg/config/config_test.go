package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Set required environment variables for testing
	os.Setenv("ENV", "development")
	os.Setenv("SERVER_PORT", "8080")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.ServerPort != 8080 {
		t.Errorf("Expected server port 8080, got %d", cfg.ServerPort)
	}

	if cfg.Environment != "development" {
		t.Errorf("Expected environment development, got %s", cfg.Environment)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_VAR", "test_value")

	value := getEnv("TEST_VAR", "default")
	if value != "test_value" {
		t.Errorf("Expected test_value, got %s", value)
	}

	value = getEnv("NON_EXISTENT_VAR", "default")
	if value != "default" {
		t.Errorf("Expected default, got %s", value)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_INT", "42")

	value := getEnvAsInt("TEST_INT", 0)
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}

	value = getEnvAsInt("NON_EXISTENT_INT", 99)
	if value != 99 {
		t.Errorf("Expected 99, got %d", value)
	}
}

func TestGetEnvAsBool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")

	value := getEnvAsBool("TEST_BOOL", false)
	if !value {
		t.Error("Expected true, got false")
	}

	value = getEnvAsBool("NON_EXISTENT_BOOL", true)
	if !value {
		t.Error("Expected true, got false")
	}
}
