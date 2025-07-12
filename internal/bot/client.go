package bot

import (
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
