package constants

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func getEnv[T string](key string) T {
	_val := os.Getenv(key)

	var val interface{} = _val

	return val.(T)
}

func getEnvDefault[T interface{}](key string, defaultValue T) T {
	_val, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	var val interface{} = _val

	return val.(T)
}

type Env struct {
	DB_NAME      string
	DB_USER      string
	DB_PASSWORD  string
	DB_HOST      string
	DB_PORT      string
	DB_SSL_MODE  string
	DB_TIME_ZONE string

	RABBITMQ_URL        string
	TRANSACTION_WORKERS int
}

func NewEnv() *Env {
	return &Env{
		DB_NAME:      getEnv("DB_NAME"),
		DB_USER:      getEnv("DB_USER"),
		DB_PASSWORD:  getEnv("DB_PASSWORD"),
		DB_HOST:      getEnv("DB_HOST"),
		DB_PORT:      getEnv("DB_PORT"),
		DB_SSL_MODE:  getEnvDefault("DB_SSL_MODE", "disable"),
		DB_TIME_ZONE: getEnvDefault("DB_TIME_ZONE", "Africa/Lagos"),

		RABBITMQ_URL: getEnv("RABBITMQ_URL"),

		TRANSACTION_WORKERS: getEnvDefault("TRANSACTION_WORKERS", 1),
	}
}
