package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBDriver string
	DBUser   string
	DBPass   string
	DBHost   string
	DBPort   string
	DBName   string
}

func LoadConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	return Config{
		Port:     getEnv("PORT", "8080"),
		DBDriver: getEnv("DB_DRIVER", "postgres"),
		DBUser:   getEnv("DB_USER", "postgres"),
		DBPass:   getEnv("DB_PASS", "postgres"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   getEnv("DB_PORT", "5432"),
		DBName:   getEnv("DB_NAME", "ecommerce"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
