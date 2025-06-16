package config

import (
	"os"
	"strconv"
)

type Config struct {
	APIPort int    // Порт для API сервера
	Env     string // Окружение (dev, prod, etc)
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	return &Config{
		APIPort: getEnvAsInt("API_PORT", 8080),
		Env:     getEnv("ENV", "development"),
	}
}

// Вспомогательная функция для чтения строковых переменных окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Вспомогательная функция для чтения числовых переменных окружения
func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return defaultValue
}