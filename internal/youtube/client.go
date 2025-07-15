package youtube

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Client handles YouTube operations
type Client struct {
	ytdlpPath string
}

// NewClient creates a new YouTube client
func NewClient() *Client {
	return &Client{
		ytdlpPath: "yt-dlp", // Assumes yt-dlp is in PATH
	}
}

// GetVideoInfo gets basic information about a YouTube video
func (c *Client) GetVideoInfo(url string) (*VideoInfo, error) {
	// Run yt-dlp to get video info as JSON
	cmd := exec.Command(c.ytdlpPath, "--print-json", "--no-download", url)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %w", err)
	}

	var videoInfo VideoInfo
	if err := json.Unmarshal(output, &videoInfo); err != nil {
		return nil, fmt.Errorf("failed to parse video info: %w", err)
	}

	return &videoInfo, nil
}

// IsValidURL performs basic URL validation
func (c *Client) IsValidURL(url string) bool {
	// Simple check - you can make this more sophisticated later
	return len(url) > 0 &&
		(strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be"))
}
