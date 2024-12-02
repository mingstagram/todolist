package config

import (
	"os"
)

type Config struct {
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	DBName string
}

func GetConfig() Config {
	return Config {
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "1111"),
		DBHost: getEnv("DB_HOST", "127.0.0.1"),
		DBPort: getEnv("DB_PORT", "3306"),
		DBName: getEnv("DB_NAME", "todolist"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}