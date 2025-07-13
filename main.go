package main

import (
	"fmt"
	"log"
	"time"

	"hamond.dev/telegram-bot-go/config"
	"hamond.dev/telegram-bot-go/internal/bot"
)

func main() {
	fmt.Println("Welcome To Hamond Bot!")
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

	// Start polling for messages
	var offset int64 = 0

	for {
		updates, err := botClient.GetUpdates(offset)
		if err != nil {
			log.Printf("Error getting updates: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, update := range updates {
			// Update offset to avoid getting the same update again
			offset = update.UpdateID + 1
			
			// Handle the message if it exists
			if update.Message != nil {
				fmt.Printf("Received message from %s: %s\n",
					update.Message.From.FirstName,
					update.Message.Text)
			
				err := botClient.HandleMessage(update.Message)
				if err != nil {
				log.Printf("Error handling message: %v", err)
				}
			}
		}

		// Small delay to avoid excessive API calls
		time.Sleep(1 * time.Second)
	}
}
