package config

import (
	"os"
	"log"

	"github.com/joho/godotenv")

type Config struct {
	TelegramBotToken string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadding .env file")
	}

	return &Config {
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}
