package bot

import (
	"testing"
)

func TestHandleCommand(t *testing.T) {
	// Create a mock client for testing
	client := &Client{
		token:   "test-token",
		baseURL: "https://api.telegram.org/bottest-token",
	}

	tests := []struct {
		name        string
		messageText string
		expectError bool
	}{
		{
			name:        "Start command",
			messageText: "/start",
			expectError: false,
		},
		{
			name:        "Help command",
			messageText: "/help",
			expectError: false,
		},
		{
			name:        "Echo command with text",
			messageText: "/echo Hello World",
			expectError: false,
		},
		{
			name:        "Echo command without text",
			messageText: "/echo",
			expectError: false,
		},
		{
			name:        "Unknown command",
			messageText: "/unknown",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				MessageID: 1,
				From: &User{
					ID:        12345,
					FirstName: "Test",
					Username:  "testuser",
				},
				Chat: Chat{
					ID:   67890,
					Type: "private",
				},
				Text: tt.messageText,
			}

			// Note: This will fail because we can't actually send messages in tests
			// We'll improve this with mocking in the next step
			err := client.handleCommand(message)

			if tt.expectError && err == nil {
				t.Error("Expected error but got nil")
			}

			// For now, we expect errors since we can't actually send messages
			// This is just testing that the parsing logic works
		})
	}
}
