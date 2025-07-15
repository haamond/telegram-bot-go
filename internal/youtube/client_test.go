package youtube

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	if client.ytdlpPath != "yt-dlp" {
		t.Errorf("Expected ytdlpPath to be 'yt-dlp', got %s", client.ytdlpPath)
	}
}

func TestIsValidURL(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "Valid YouTube URL",
			url:      "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "Valid YouTube short URL",
			url:      "https://youtu.be/dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "Invalid URL - empty",
			url:      "",
			expected: false,
		},
		{
			name:     "Invalid URL - not YouTube",
			url:      "https://www.google.com",
			expected: false,
		},
		{
			name:     "Invalid URL - random text",
			url:      "not a url at all",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.IsValidURL(tt.url)
			if result != tt.expected {
				t.Errorf("IsValidURL(%s) = %v, expected %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestGetVideoInfo(t *testing.T) {
	client := NewClient()

	// Test with a known working URL
	url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	info, err := client.GetVideoInfo(url)
	if err != nil {
		t.Fatalf("GetVideoInfo() failed: %v", err)
	}

	if info == nil {
		t.Fatal("GetVideoInfo() returned nil info")
	}

	// Test that we got basic fields
	if info.ID == "" {
		t.Error("Expected non-empty ID")
	}

	if info.Title == "" {
		t.Error("Expected non-empty Title")
	}

	if info.Duration <= 0 {
		t.Error("Expected positive Duration")
	}

	if info.Uploader == "" {
		t.Error("Expected non-empty Uploader")
	}

	// Test specific values for this well-known video
	if info.ID != "dQw4w9WgXcQ" {
		t.Errorf("Expected ID 'dQw4w9WgXcQ', got %s", info.ID)
	}
}

func TestGetVideoInfoInvalidURL(t *testing.T) {
	client := NewClient()

	// Test with invalid URL
	_, err := client.GetVideoInfo("https://www.google.com")
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}
