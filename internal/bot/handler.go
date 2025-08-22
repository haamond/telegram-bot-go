package bot

import (
	"fmt"
	"os"
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

	// Check if the message contains a YouTube URL
	if c.youtube.IsValidURL(message.Text) {
		return c.handleDownloadCommand(message.Chat.ID, strings.TrimSpace(message.Text))
	}

	// For non-YouTube URLs, provide help
	return c.SendMessage(message.Chat.ID, "👋 Send me a YouTube link and I'll download the video for you!\n\nExample: https://youtube.com/watch?v=...\n\nOr use /help to see available commands.")
}

// handleCommand processes bot commands
func (c *Client) handleCommand(message *Message) error {
	command := strings.ToLower(message.Text)

	switch {
	case strings.HasPrefix(command, "/start"):
		welcomeText := fmt.Sprintf("Hello %s! 👋\n\nI'm your YouTube downloader bot. Just send me a YouTube link and I'll download the video for you!\n\n📹 Supported formats:\n• YouTube URLs (youtube.com/watch?v=...)\n• YouTube short URLs (youtu.be/...)\n\nThe video will be downloaded in 360p quality for optimal file size and compatibility.\n\nType /help for more info.", message.From.FirstName)
		return c.SendMessage(message.Chat.ID, welcomeText)

	case strings.HasPrefix(command, "/help"):
		helpText := "📖 *How to use this bot:*\n\n1️⃣ Send me any YouTube link\n2️⃣ I'll download the video (360p)\n3️⃣ The video will be sent back to you\n\n*Commands:*\n/start - Welcome message\n/help - This help message\n/download <url> - Explicitly download a video\n\n*Examples:*\n• https://youtube.com/watch?v=dQw4w9WgXcQ\n• https://youtu.be/dQw4w9WgXcQ\n\n⚡ Just paste the link and I'll handle the rest!"
		return c.SendMessage(message.Chat.ID, helpText)

	case strings.HasPrefix(command, "/download "):
		url := strings.TrimPrefix(message.Text, "/download ")
		url = strings.TrimSpace(url)
		if url == "" {
			return c.SendMessage(message.Chat.ID, "Please provide a YouTube URL. Example: /download https://youtube.com/watch?v=...")
		}
		return c.handleDownloadCommand(message.Chat.ID, url)

	default:
		return c.SendMessage(message.Chat.ID, "❓ Unknown command. Type /help to see available commands.")
	}
}

// handleDownloadCommand handles video download requests
func (c *Client) handleDownloadCommand(chatID int64, url string) error {
	// Clean the URL (remove any extra spaces or characters)
	url = strings.TrimSpace(url)

	// Validate URL
	if !c.youtube.IsValidURL(url) {
		return c.SendMessage(chatID, "❌ Invalid YouTube URL. Please provide a valid YouTube or youtu.be link.")
	}

	// Send "processing" message
	err := c.SendMessage(chatID, "🔍 Fetching video information...")
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

	// Format duration nicely
	duration := formatDuration(videoInfo.Duration)

	// Send download starting message with more info
	err = c.SendMessage(chatID, fmt.Sprintf("📹 *%s*\n\n⏱ Duration: %s\n📊 Quality: 360p\n\n⬇️ Downloading video...",
		videoInfo.Title, duration))
	if err != nil {
		return err
	}

	// Create a simple filename
	filename := fmt.Sprintf("%s.%%(ext)s", videoInfo.ID) // yt-dlp will replace %%(ext)s with actual extension

	// Download the video
	fmt.Printf("Starting download: %s\n", videoInfo.Title)
	err = c.youtube.DownloadFormat18(url, filename)
	if err != nil {
		fmt.Printf("Download failed %v\n", err)
		return c.SendMessage(chatID, "❌ Download failed. This might be due to:\n• Video is private or age-restricted\n• Video is too long\n• Regional restrictions\n\nPlease try another video.")
	}

	// Find the actual downloaded file (yt-dlp replaces %%(ext)s with actual extension)
	downloadedFile := fmt.Sprintf("%s.mp4", videoInfo.ID) // Format 18 is always mp4
	fmt.Printf("Download completed: %s -> %s\n", videoInfo.Title, downloadedFile)

	// Check file size before uploading (Telegram has a 50MB limit for bots)
	fileInfo, err := os.Stat(downloadedFile)
	if err == nil && fileInfo.Size() > 50*1024*1024 {
		os.Remove(downloadedFile)
		return c.SendMessage(chatID, "❌ Video is too large (>50MB). Telegram bots can only send files up to 50MB.\n\nTry a shorter video.")
	}

	// Send upload message
	err = c.SendMessage(chatID, "📤 Uploading to Telegram...")
	if err != nil {
		return err
	}

	// Send the video file back to user
	fmt.Printf("Uploading file to Telegram: %s\n", downloadedFile)
	err = c.SendVideo(chatID, downloadedFile)
	if err != nil {
		fmt.Printf("Upload failed: %v\n", err)
		os.Remove(downloadedFile)
		return c.SendMessage(chatID, "❌ Failed to upload video to Telegram. The file might be too large or in an unsupported format.")
	}

	// Clean up downloaded file
	fmt.Printf("Cleaning up file: %s\n", downloadedFile)
	os.Remove(downloadedFile)

	fmt.Printf("Process completed successfully for: %s\n", videoInfo.Title)
	return c.SendMessage(chatID, "✅ Video sent successfully! Send another link to download more videos.")
}

// formatDuration converts seconds to a human-readable format
func formatDuration(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%d seconds", seconds)
	}
	minutes := seconds / 60
	secs := seconds % 60
	if minutes < 60 {
		if secs == 0 {
			return fmt.Sprintf("%d minutes", minutes)
		}
		return fmt.Sprintf("%d min %d sec", minutes, secs)
	}
	hours := minutes / 60
	mins := minutes % 60
	return fmt.Sprintf("%d hr %d min", hours, mins)
}
