package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hamond.dev/telegram-bot-go/config"
	"hamond.dev/telegram-bot-go/internal/bot"
)

func main() {
	cfg := config.Load()

	if cfg.TelegramBotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	// Create bot client
	botClient := bot.NewClient(cfg.TelegramBotToken)

	// Test the connection
	user, err := botClient.GetMe()
	if err != nil {
		log.Fatalf("Failed to get bot info: %v", err)
	}

	fmt.Printf("Bot started successfully!\n")
	fmt.Printf("Bot Name: %s\n", user.FirstName)
	fmt.Printf("Bot Username: @%s\n", user.Username)
	fmt.Printf("Mode: %s\n", cfg.Mode)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start bot based on mode
	switch cfg.Mode {
	case "webhook":
		startWebhookMode(botClient, cfg, sigChan)
	case "polling":
		startPollingMode(botClient, sigChan)
	default:
		log.Fatalf("Invalid mode: %s. Use 'polling' or 'webhook'", cfg.Mode)
	}
}

func startPollingMode(botClient *bot.Client, sigChan chan os.Signal) {
	fmt.Println("Starting in polling mode...")
	fmt.Println("Waiting for messages... (Press Ctrl+C to stop)")

	var offset int64 = 0

	// Start polling in a goroutine
	go func() {
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
	}()

	// Wait for shutdown signal
	<-sigChan
	fmt.Println("\nShutting down bot...")
}

func startWebhookMode(botClient *bot.Client, cfg *config.Config, sigChan chan os.Signal) {
	if cfg.WebhookURL == "" {
		log.Fatal("WEBHOOK_URL must be set for webhook mode")
	}

	fmt.Printf("Starting in webhook mode...\n")
	fmt.Printf("Webhook URL: %s\n", cfg.WebhookURL)

	// Delete any existing webhook first
	err := botClient.DeleteWebhook()
	if err != nil {
		log.Printf("Warning: Failed to delete existing webhook: %v", err)
	}

	// Set the new webhook
	webhookEndpoint := cfg.WebhookURL + "/webhook"
	err = botClient.SetWebhook(webhookEndpoint)
	if err != nil {
		log.Fatalf("Failed to set webhook: %v", err)
	}

	fmt.Printf("Webhook set successfully to: %s\n", webhookEndpoint)

	// Create and start the server
	server := bot.NewServer(botClient, cfg.Port)

	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	fmt.Printf("Server started on port %s\n", cfg.Port)
	fmt.Println("Waiting for webhook updates... (Press Ctrl+C to stop)")

	// Wait for shutdown signal
	<-sigChan
	fmt.Println("\nShutting down webhook...")

	// Clean up webhook on shutdown
	err = botClient.DeleteWebhook()
	if err != nil {
		log.Printf("Warning: Failed to delete webhook on shutdown: %v", err)
	} else {
		fmt.Println("Webhook deleted successfully")
	}
}
