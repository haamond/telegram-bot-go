# Telegram Bot in Go

A learning project to build a Telegram bot using Go, following best practices.

## Features

- [x] Environment configuration with .env support
- [x] Bot client with API communication
- [x] GetMe API endpoint implementation
- [x] Message handling with polling
- [x] Command processing (/start, /help, /echo)
- [x] Echo functionality for regular messages
- [ ] Webhook support (coming next)
- [ ] Advanced command features
- [ ] Database integration

## Current Commands

- `/start` - Welcome message with bot introduction
- `/help` - Display available commands
- `/echo <text>` - Echo back the provided text
- Any regular message will be echoed back

## Setup

1. Clone the repository:
  ```bash
  git clone https://github.com/haamond/telegram-bot-go.git
  cd telegram-bot-go

2. Install dependencies:  
3. Create a .env file with your bot token:  
4. Run the bot:  
5. Find your bot on Telegram and send /start to begin!  

## Project Structure
telegram-bot-go/
├── config/          # Configuration management
├── internal/bot/    # Bot implementation
│   ├── client.go    # HTTP client and API methods
│   ├── handler.go   # Message and command handling
│   └── types.go     # Telegram API type definitions
├── .env            # Environment variables (not committed)
├── main.go         # Application entry point with polling loop
└── go.mod          # Go module definition




