package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	AppName                  string
	Host                     string
	DB_HOST                  string
	DB_USER                  string
	DB_PASSWORD              string
	DB_NAME                  string
	DB_PORT                  string
	TIME_ZONE                string
	SERVER_HOST              string
	SERVER_PORT              string
	AUTO_MIGRATE             string
	ADMIN_EMAIL              string
	ADMIN_PASSWORD           string
	ACCESS_TOKEN_EXPIRATION  string
	REFRESH_TOKEN_EXPIRATION string
	JWT_SECRET_KEY           string
)

func LoadEnv() {
	wd, _ := os.Getwd()
	root := filepath.Join(wd, "../../apps/server/src/")
	envPath := filepath.Join(root, "config", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("⚠️  No .env file loaded from %s: %v", envPath, err)
	} else {
		log.Printf("✅ Loaded .env file from %s", envPath)
	}

	AppName = getEnv("APP_NAME", "SodNix")
	DB_HOST = getEnv("DB_HOST", "localhost")
	DB_USER = getEnv("DB_USER", "postgres")
	DB_PASSWORD = getEnv("DB_PASSWORD", "123456")
	DB_NAME = getEnv("DB_NAME", "src")
	DB_PORT = getEnv("DB_PORT", "5432")
	TIME_ZONE = getEnv("TIME_ZONE", "Asia/Ho_Chi_Minh")
	SERVER_HOST = getEnv("SERVER_HOST", "127.0.0.1")
	SERVER_PORT = getEnv("SERVER_PORT", "8080")
	AUTO_MIGRATE = getEnv("AUTO_MIGRATE", "false")
	ADMIN_EMAIL = getEnv("ADMIN_EMAIL", "admin@example.com")
	ADMIN_PASSWORD = getEnv("ADMIN_PASSWORD", "admin")
	ACCESS_TOKEN_EXPIRATION = getEnv("ACCESS_TOKEN_EXPIRATION", "15")
	REFRESH_TOKEN_EXPIRATION = getEnv("REFRESH_TOKEN_EXPIRATION", "7")
	JWT_SECRET_KEY = getEnv("JWT_SECRET_KEY", "your_secret_key")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
