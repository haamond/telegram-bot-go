package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	WebhookURL       string
	Port             string
	Mode             string // "polling" or "webhook"
}

func Load() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Default values
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "polling" // Default to polling
	}

	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		WebhookURL:       os.Getenv("WEBHOOK_URL"),
		Port:             port,
		Mode:             mode,
	}
}
