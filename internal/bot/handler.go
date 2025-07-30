package bot

import (
	"fmt"
	"strings"
)

// HandleMessage processes incoming messages
func (c *Client) HandleMessage(message *Message) error {
	if message.Text == "" {
		return nil // Ignore non-text messages for now
	}

	// Handle commands (messages starting with /)
	if strings.HasPrefix(message.Text, "/") {
		return c.handleCommand(message)
	}

	// Echo back regular messages
	return c.SendMessage(message.Chat.ID, "You said: "+message.Text)
}

// handleCommand processes bot commands
func (c *Client) handleCommand(message *Message) error {
	command := strings.ToLower(message.Text)

	switch {
	case strings.HasPrefix(command, "/start"):
		welcomeText := fmt.Sprintf("Hello %s! 👋\n\nI'm your Go YouTube downloader bot. Try these commands:\n/help - Show available commands\n/download <youtube_url> - Download video in 360p", message.From.FirstName)
		return c.SendMessage(message.Chat.ID, welcomeText)

	case strings.HasPrefix(command, "/help"):
		helpText := "Available commands:\n/start - Welcome message\n/help - This help message\n/echo <text> - Echo your message\n/download <youtube_url> - Download video in 360p\n\nOr just send me any message and I'll echo it back!"
		return c.SendMessage(message.Chat.ID, helpText)

	case strings.HasPrefix(command, "/echo "):
		echoText := strings.TrimPrefix(message.Text, "/echo ")
		if echoText == "" {
			return c.SendMessage(message.Chat.ID, "Please provide text to echo. Example: /echo Hello World")
		}
		return c.SendMessage(message.Chat.ID, "Echo: "+echoText)

	case strings.HasPrefix(command, "/download "):
		url := strings.TrimPrefix(message.Text, "/download ")
		url = strings.TrimSpace(url)
		if url == "" {
			return c.SendMessage(message.Chat.ID, "Please provide a YouTube URL. Example: /download https://youtube.com/watch?v=...")
		}
		return c.handleDownloadCommand(message.Chat.ID, url)

	default:
		return c.SendMessage(message.Chat.ID, "Unknown command. Type /help to see available commands.")
	}
}

// handleDownloadCommand handles the /download command (simplified to format 18)
func (c *Client) handleDownloadCommand(chatID int64, url string) error {
	// Validate URL
	if !c.youtube.IsValidURL(url) {
		return c.SendMessage(chatID, "❌ Invalid YouTube URL. Please provide a valid YouTube or youtu.be link.")
	}

	// Send "processing" message
	err := c.SendMessage(chatID, "🔍 Getting video information...")
	if err != nil {
		return err
	}

	// Get video info
	videoInfo, err := c.youtube.GetVideoInfo(url)
	if err != nil {
		fmt.Printf("Error getting video info: %v\n", err)
		return c.SendMessage(chatID, "❌ Failed to get video information. Please check the URL and try again.")
	}

	fmt.Printf("Video info: Title=%s, Duration=%d seconds\n", videoInfo.Title, videoInfo.Duration)

	return c.SendMessage(chatID, fmt.Sprintf("📹 *%s*\n\nDuration: %d seconds\n\nDownloading in 360p (format 18)...",
		videoInfo.Title, videoInfo.Duration))
}
