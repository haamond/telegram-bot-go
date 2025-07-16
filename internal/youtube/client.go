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

// GetAvailableFormats gets all available formats for a video
func (c *Client) GetAvailableFormats(url string) ([]VideoFormat, error) {
	// First, let's try getting video info with format details
	cmd := exec.Command(c.ytdlpPath, "--dump-json", "--no-download", url)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %w", err)
	}

	// Parse the JSON response with flexible quality field
	var videoData struct {
		Formats []map[string]interface{} `json:"formats"`
	}

	if err := json.Unmarshal(output, &videoData); err != nil {
		return nil, fmt.Errorf("failed to parse video data: %w", err)
	}

	var formats []VideoFormat
	for _, f := range videoData.Formats {
		format := VideoFormat{}

		// Parse FormatID
		if id, ok := f["format_id"].(string); ok {
			format.FormatID = id
		}

		// Parse Extension
		if ext, ok := f["ext"].(string); ok {
			format.Extension = ext
		}

		// Parse FileSize
		if size, ok := f["filesize"].(float64); ok {
			format.FileSize = int64(size)
		}

		// Parse Width
		if width, ok := f["width"].(float64); ok {
			format.Width = int(width)
		}

		// Parse Height
		if height, ok := f["height"].(float64); ok {
			format.Height = int(height)
		}

		// Parse FPS
		if fps, ok := f["fps"].(float64); ok {
			format.FPS = fps
		}

		// Parse codecs
		vcodec, hasVCodec := f["vcodec"].(string)
		acodec, hasACodec := f["acodec"].(string)

		format.HasVideo = hasVCodec && vcodec != "none" && vcodec != ""
		format.HasAudio = hasACodec && acodec != "none" && acodec != ""

		// Parse Quality (can be string or number)
		if quality, ok := f["quality"].(string); ok {
			format.Quality = quality
		} else if quality, ok := f["quality"].(float64); ok {
			format.Quality = fmt.Sprintf("%.0f", quality)
		} else {
			// Determine quality from height or format_note
			if format.Height > 0 {
				format.Quality = fmt.Sprintf("%dp", format.Height)
			} else if formatNote, ok := f["format_note"].(string); ok && formatNote != "" {
				format.Quality = formatNote
			} else {
				format.Quality = "unknown"
			}
		}

		formats = append(formats, format)
	}

	return formats, nil
}
