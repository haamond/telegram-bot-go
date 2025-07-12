package main

import (
	"fmt"
	"log"

	"hamond.dev/telegram-bot-go/config"
	"hamond.dev/telegram-bot-go/internal/bot"
)

func main() {
	fmt.Println("Hello, Mom!")
	cfg := config.Load()

	if cfg.TelegramBotToken == "" {
		fmt.Println("TELEGRAM_BOT_TOKEN is not set")
		return
	}

	// Create bot client
	botClient := bot.NewClient(cfg.TelegramBotToken)

	// Test the connection
	user, err := botClient.GetMe()
	if err != nil {
		log.Fatalf("Failed to get bot info: %v", err)
	}

	fmt.Printf("Bot connected successfully!\n")
	fmt.Printf("Bot Name: %s\n", user.FirstName)
    	fmt.Printf("Bot Username: @%s\n", user.Username)
    	fmt.Printf("Bot ID: %d\n", user.ID)

}
