package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	token   string
	baseURL string
}

// NewClient creates a new bot client
func NewClient(token string) *Client {
	return &Client{
		token:   token,
		baseURL: "https://api.telegram.org/bot" + token,
	}
}

// GetMe returns basic information about the bot
func (c *Client) GetMe() (*User, error) {
	url := c.baseURL + "/getMe"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response GetMeResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !response.Ok {
		return nil, fmt.Errorf("API error: %s", response.Description)
	}

	return &response.Result, nil
}

// GetUpdates retrieves new updates from Telegram
func (c *Client) GetUpdates(offset int64) ([]Update, error) {
	url := fmt.Sprintf("%s/getUpdates?offset=%d", c.baseURL, offset)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get updates: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response GetUpdatesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !response.Ok {
		return nil, fmt.Errorf("API error getting updates")
	}

	return response.Result, nil
}

// SendMessage sends a text message to a chat
func (c *Client) SendMessage(chatID int64, text string) error {
	requestBody := SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	url := c.baseURL + "/sendMessage"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var response SendMessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !response.Ok {
		return fmt.Errorf("API error sending message")
	}

	return nil
}

// SetWebhook sets the webhook URL for the bot
func (c *Client) SetWebhook(webhookURL string) error {
	requestBody := SetWebhookRequest{
		URL: webhookURL,
	}

	requestJsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook request: %w", err)
	}

	url := c.baseURL + "/setWebhook"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestJsonBody))
	if err != nil {
		return fmt.Errorf("failed to set webhook: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var response SetWebhookResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !response.Ok {
		return fmt.Errorf("API error setting webhook: %s", response.Description)
	}

	return nil
}

// DeleteWebhook removes the webhook (returns to polling mode)
func (c *Client) DeleteWebhook() error {
	url := c.baseURL + "/deleteWebhook"

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

// GetWebhookInfo gets current webhook information
func (c *Client) GetWebhookInfo() (*WebhookInfo, error) {
	url := c.baseURL + "/getWebhookInfo"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response GetWebhookInfoResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !response.Ok {
		return nil, fmt.Errorf("API error getting webhook info")
	}

	return &response.Result, nil
}
