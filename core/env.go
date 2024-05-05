package core

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func getEnv(key string) string {
	return os.Getenv(key)
}

func getEnvDefault(key string, defaultValue string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return val
}

var (
	DB_NAME      = getEnv("DB_NAME")
	DB_USER      = getEnv("DB_USER")
	DB_PASSWORD  = getEnv("DB_PASSWORD")
	DB_HOST      = getEnv("DB_HOST")
	DB_PORT      = getEnv("DB_PORT")
	DB_SSL_MODE  = getEnvDefault("DB_SSL_MODE", "disable")
	DB_TIME_ZONE = getEnvDefault("DB_TIME_ZONE", "disable")
)
