package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	SMTPHost     string
	SMTPPort     int
	SMTPFrom     string
	SMTPPassword string
	IsProduction bool
	LogFile      string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	smtpPort, err := strconv.Atoi(mustGetEnv("SMTP_PORT"))

	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	return &Config{
		SMTPPort:     smtpPort,
		SMTPHost:     mustGetEnv("SMTP_HOST"),
		SMTPFrom:     mustGetEnv("SMTP_FROM"),
		SMTPPassword: mustGetEnv("SMTP_PASSWORD"),
		Port:         mustGetEnv("PORT"),
		IsProduction: mustGetEnv("IS_PROD") == "true",
		LogFile:      mustGetEnv("LOG_FILE"),
	}
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Required env variable %s is not set. Value: %v", key, val)
	}

	return val
}
