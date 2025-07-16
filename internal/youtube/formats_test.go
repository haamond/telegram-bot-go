package youtube

import (
	"testing"
)

func TestGetAvailableFormats(t *testing.T) {
	client := NewClient()

	// Test with a known working URL
	url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	formats, err := client.GetAvailableFormats(url)
	if err != nil {
		t.Fatalf("GetAvailableFormats() failed: %v", err)
	}

	if len(formats) == 0 {
		t.Error("Expected at least one format, got 0")
	}

	// Log first few formats for debugging
	t.Logf("Found %d formats", len(formats))
	for i, format := range formats {
		if i >= 5 { // Log first 5 formats
			break
		}
		t.Logf("Format %d: ID=%s, Quality=%s, Ext=%s, HasVideo=%v, HasAudio=%v, Size=%d",
			i, format.FormatID, format.Quality, format.Extension, format.HasVideo, format.HasAudio, format.FileSize)
	}

	// Check that we have at least some formats with required fields
	validFormats := 0
	for _, format := range formats {
		if format.FormatID != "" && format.Extension != "" {
			validFormats++
		}
	}

	if validFormats == 0 {
		t.Error("Expected at least one format with valid FormatID and Extension")
	}

	// Check that we have at least some video formats
	videoFormats := 0
	for _, format := range formats {
		if format.HasVideo {
			videoFormats++
		}
	}

	if videoFormats == 0 {
		t.Error("Expected at least one video format")
	}

	// Check that we have some mobile-friendly formats
	mobileFormats := FilterMobileFriendlyFormats(formats)
	t.Logf("Found %d mobile-friendly formats", len(mobileFormats))

	for i, format := range mobileFormats {
		if i >= 3 { // Log first 3 mobile formats
			break
		}
		t.Logf("Mobile Format %d: Quality=%s, Ext=%s, Size=%s",
			i, format.Quality, format.Extension, FormatSizeToString(format.FileSize))
	}
}

func TestFilterMobileFriendlyFormats(t *testing.T) {
	// Test data - simulating yt-dlp output
	allFormats := []VideoFormat{
		{FormatID: "18", Quality: "360p", Extension: "mp4", FileSize: 50000000, HasVideo: true, HasAudio: true},
		{FormatID: "22", Quality: "720p", Extension: "mp4", FileSize: 150000000, HasVideo: true, HasAudio: true},
		{FormatID: "137", Quality: "1080p", Extension: "mp4", FileSize: 300000000, HasVideo: true, HasAudio: false}, // Video only
		{FormatID: "140", Quality: "audio", Extension: "m4a", FileSize: 20000000, HasVideo: false, HasAudio: true},  // Audio only
		{FormatID: "251", Quality: "audio", Extension: "webm", FileSize: 25000000, HasVideo: false, HasAudio: true}, // Audio only
	}

	mobileFormats := FilterMobileFriendlyFormats(allFormats)

	// Should only include formats with both video and audio
	for _, format := range mobileFormats {
		if !format.HasVideo {
			t.Errorf("Mobile format should have video: %+v", format)
		}
		if !format.HasAudio {
			t.Errorf("Mobile format should have audio: %+v", format)
		}
		if format.Extension != "mp4" {
			t.Errorf("Expected mp4 format for mobile, got %s", format.Extension)
		}
	}

	// Should have at least 360p and 720p
	hasLowRes := false
	hasHighRes := false
	for _, format := range mobileFormats {
		if format.Quality == "360p" {
			hasLowRes = true
		}
		if format.Quality == "720p" {
			hasHighRes = true
		}
	}

	if !hasLowRes {
		t.Error("Expected at least one low resolution format (360p)")
	}
	if !hasHighRes {
		t.Error("Expected at least one high resolution format (720p)")
	}
}

func TestFormatSizeToString(t *testing.T) {
	tests := []struct {
		name     string
		bytes    int64
		expected string
	}{
		{
			name:     "Bytes",
			bytes:    500,
			expected: "500 B",
		},
		{
			name:     "Kilobytes",
			bytes:    1500,
			expected: "1.5 KB",
		},
		{
			name:     "Megabytes",
			bytes:    50000000,
			expected: "50.0 MB", // Using decimal: 50,000,000 / 1,000,000 = 50.0
		},
		{
			name:     "Gigabytes",
			bytes:    1500000000,
			expected: "1.5 GB", // Using decimal: 1,500,000,000 / 1,000,000,000 = 1.5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSizeToString(tt.bytes)
			if result != tt.expected {
				t.Errorf("FormatSizeToString(%d) = %s, expected %s", tt.bytes, result, tt.expected)
			}
		})
	}
}

func TestCreateFormatMessage(t *testing.T) {
	formats := []VideoFormat{
		{FormatID: "18", Quality: "360p", Extension: "mp4", FileSize: 50000000, HasVideo: true, HasAudio: true},
		{FormatID: "22", Quality: "720p", Extension: "mp4", FileSize: 150000000, HasVideo: true, HasAudio: true},
	}

	message := CreateFormatMessage("Test Video", formats)

	if message == "" {
		t.Error("Expected non-empty format message")
	}

	// Check that message contains expected information
	expectedContent := []string{"Test Video", "360p", "720p", "50.0 MB", "150.0 MB"}
	for _, content := range expectedContent {
		if !contains(message, content) {
			t.Errorf("Expected message to contain '%s', but it didn't. Message: %s", content, message)
		}
	}
}

// Helper function for string contains check
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) &&
			(hasSubstring(s, substr))))
}

func hasSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
