# YouTube Downloader Telegram Bot

A simple and efficient Telegram bot built with Go that downloads YouTube videos. Just send a YouTube link and get the video back!

## ✨ Features

- **Instant Download**: Just paste a YouTube link - no commands needed!
- **Smart Recognition**: Automatically detects YouTube URLs in messages
- **360p Quality**: Optimized for quick downloads and Telegram's file size limits
- **User-Friendly**: Simple interface with helpful messages
- **Multiple Modes**: Supports both polling and webhook modes
- **Clean Architecture**: Well-structured Go code following best practices

## 🚀 Quick Start

### Prerequisites

- Go 1.24+ installed
- [yt-dlp](https://github.com/yt-dlp/yt-dlp) installed (`pip install yt-dlp` or `brew install yt-dlp`)
- A Telegram Bot Token from [@BotFather](https://t.me/botfather)

### Installation

1. **Clone the repository:**
```bash
git clone https://github.com/yourusername/telegram-yt-bot.git
cd telegram-yt-bot
```

2. **Install dependencies:**
```bash
go mod download
```

3. **Set up environment variables:**
Create a `.env` file in the project root:
```env
TELEGRAM_BOT_TOKEN=your_bot_token_here
MODE=polling
# Optional for webhook mode:
# WEBHOOK_URL=https://your-domain.com
# PORT=8080
```

4. **Run the bot:**
```bash
go run main.go
```

## ⚙️ Configuration

The bot supports two modes of operation:

### Polling Mode (Default)
- Simple setup, no external dependencies
- Bot polls Telegram servers for updates
- Perfect for development and small-scale usage

```env
MODE=polling
TELEGRAM_BOT_TOKEN=your_bot_token
```

### Webhook Mode
- More efficient for production
- Requires a public URL (HTTPS)
- Telegram pushes updates to your server

```env
MODE=webhook
TELEGRAM_BOT_TOKEN=your_bot_token
WEBHOOK_URL=https://your-domain.com
PORT=8080
```

## 📱 Usage

1. **Start a chat** with your bot on Telegram
2. **Send any YouTube URL** - the bot will automatically detect it
3. **Wait for download** - the bot will process and send you the video
4. **Enjoy!** Videos are downloaded in 360p for optimal size and speed

### Supported URL Formats
- `https://www.youtube.com/watch?v=VIDEO_ID`
- `https://youtu.be/VIDEO_ID`
- `https://m.youtube.com/watch?v=VIDEO_ID`

## 🏗️ Project Structure

```
telegram-yt-bot/
├── main.go              # Application entry point
├── config/              # Configuration management
│   ├── config.go        # Config struct and loader
│   └── config_test.go   # Config tests
├── internal/bot/        # Bot implementation
│   ├── client.go        # Telegram API client
│   ├── handler.go       # Message handling logic
│   ├── server.go        # Webhook server
│   └── types.go         # Bot type definitions
├── internal/youtube/    # YouTube integration
│   ├── client.go        # YouTube downloader client
│   ├── formats.go       # Video format handling
│   └── types.go         # YouTube type definitions
└── internal/storage/    # Future storage implementation
```

## 🧪 Testing

Run the test suite:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

## 🚀 Deployment

### Docker (Coming Soon)
```bash
docker build -t telegram-yt-bot .
docker run -e TELEGRAM_BOT_TOKEN=your_token telegram-yt-bot
```

### Heroku
1. Create a Heroku app
2. Set environment variables in Heroku dashboard
3. Deploy using Git or GitHub integration

### VPS/Cloud Server
1. Copy the binary to your server
2. Set up environment variables
3. Run as a systemd service for production

## 🛠️ Development

### Prerequisites
- Go 1.24+
- yt-dlp installed locally
- A Telegram bot token for testing

### Local Development
```bash
# Clone and setup
git clone <repository>
cd telegram-yt-bot
go mod download

# Create .env file with your test bot token
echo "TELEGRAM_BOT_TOKEN=your_test_token" > .env
echo "MODE=polling" >> .env

# Run in development
go run main.go
```

## 📚 API Documentation

The bot uses the Telegram Bot API and yt-dlp for YouTube downloads:

- **Telegram Bot API**: [Official Documentation](https://core.telegram.org/bots/api)
- **yt-dlp**: [GitHub Repository](https://github.com/yt-dlp/yt-dlp)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This bot is for educational purposes. Please respect YouTube's Terms of Service and only download content you have permission to download.

## 🐛 Troubleshooting

### Common Issues

**Bot not responding:**
- Check your bot token is correct
- Ensure the bot is not already running elsewhere
- Verify network connectivity

**Download fails:**
- Update yt-dlp: `pip install --upgrade yt-dlp`
- Check if the YouTube video is available
- Some videos may be geo-restricted

**Webhook issues:**
- Ensure your webhook URL is accessible via HTTPS
- Check firewall settings
- Verify the webhook endpoint is correct

### Getting Help

- Check the [Issues](https://github.com/yourusername/telegram-yt-bot/issues) page
- Create a new issue with detailed information
- Include logs and error messages when reporting bugs
