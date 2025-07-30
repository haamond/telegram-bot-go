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

// DownloadFormat18 downloads a video using format 18 (360p mp4 with audio)
func (c *Client) DownloadFormat18(url, outputPath string) error {
	// Run yt-dlp to download format 18
	cmd := exec.Command(c.ytdlpPath, "-f", "18", "-o", outputPath, url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download video: %w (output: %s)", err, string(output))
	}

	return nil
}
