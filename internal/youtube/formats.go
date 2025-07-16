package youtube

import (
	"fmt"
	"sort"
	"strings"
)

// FilterMobileFriendlyFormats filters formats suitable for mobile devices
func FilterMobileFriendlyFormats(formats []VideoFormat) []VideoFormat {
	var mobileFormats []VideoFormat

	// Define mobile-friendly qualities
	mobileQualities := map[string]int{
		"360p":  360,
		"480p":  480,
		"720p":  720,
		"1080p": 1080,
	}

	for _, format := range formats {
		// Only include formats with both video and audio
		if !format.HasVideo || !format.HasAudio {
			continue
		}

		// Only include mp4 format - best mobile compatibility
		if format.Extension != "mp4" {
			continue
		}

		// Only include mobile-friendly qualities
		if _, isMobileFriendly := mobileQualities[format.Quality]; isMobileFriendly {
			mobileFormats = append(mobileFormats, format)
		}
	}

	// Sort by quality (ascending)
	sort.Slice(mobileFormats, func(i, j int) bool {
		qi := mobileQualities[mobileFormats[i].Quality]
		qj := mobileQualities[mobileFormats[j].Quality]
		return qi < qj
	})

	return mobileFormats
}

// FormatSizeToString converts bytes to human-readable format using decimal units
func FormatSizeToString(bytes int64) string {
	const (
		KB = 1000      // Changed from 1024 to 1000 (decimal)
		MB = KB * 1000 // Changed from 1024 to 1000 (decimal)
		GB = MB * 1000 // Changed from 1024 to 1000 (decimal)
	)

	if bytes < KB {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < MB {
		return fmt.Sprintf("%.1f KB", float64(bytes)/KB)
	} else if bytes < GB {
		return fmt.Sprintf("%.1f MB", float64(bytes)/MB)
	} else {
		return fmt.Sprintf("%.1f GB", float64(bytes)/GB)
	}
}

// CreateFormatMessage creates a user-friendly message showing available formats
func CreateFormatMessage(videoTitle string, formats []VideoFormat) string {
	if len(formats) == 0 {
		return "No mobile-friendly formats available for this video."
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("📱 *%s*\n\n", videoTitle))
	builder.WriteString("Available mobile-friendly formats:\n\n")

	for i, format := range formats {
		sizeStr := "Unknown size"
		if format.FileSize > 0 {
			sizeStr = FormatSizeToString(format.FileSize)
		}

		builder.WriteString(fmt.Sprintf("%d. *%s* - %s (%s)\n",
			i+1, format.Quality, sizeStr, format.Extension))
	}

	builder.WriteString("\nSend the number of your preferred quality to download!")

	return builder.String()
}
