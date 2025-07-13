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
        welcomeText := fmt.Sprintf("Hello %s! 👋\n\nI'm your bot. Try these commands:\n/help - Show available commands\n/echo <text> - Echo your message", message.From.FirstName)
        return c.SendMessage(message.Chat.ID, welcomeText)
        
    case strings.HasPrefix(command, "/help"):
        helpText := "Available commands:\n/start - Welcome message\n/help - This help message\n/echo <text> - Echo your message\n\nOr just send me any message and I'll echo it back!"
        return c.SendMessage(message.Chat.ID, helpText)
        
    case strings.HasPrefix(command, "/echo "):
        echoText := strings.TrimPrefix(message.Text, "/echo ")
        if echoText == "" {
            return c.SendMessage(message.Chat.ID, "Please provide text to echo. Example: /echo Hello World")
        }
        return c.SendMessage(message.Chat.ID, "Echo: "+echoText)
        
    default:
        return c.SendMessage(message.Chat.ID, "Unknown command. Type /help to see available commands.")
    }
}
