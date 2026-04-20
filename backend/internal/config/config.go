package config

import "os"

const (
	defaultPort     = "8080"
	defaultLogLevel = "info"
)

type Config struct {
	Port     string
	LogLevel string
}

func Load() Config {
	return Config{
		Port:     getEnv("PORT", defaultPort),
		LogLevel: getEnv("LOG_LEVEL", defaultLogLevel),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}