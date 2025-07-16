package youtube

// VideoInfo represents basic video information
type VideoInfo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	Uploader    string `json:"uploader"`
	Description string `json:"description"`
	URL         string `json:"webpage_url"`
}

// DownloadRequest represents a download request
type DownloadRequest struct {
	URL    string
	ChatID int64
	Format string // Will be something like "720p", "480p", etc.
}

// DownloadResult represents the result of a download
type DownloadResult struct {
	Success  bool
	FilePath string
	Error    string
}

// VideoFormat represents a video format/quality option
type VideoFormat struct {
	FormatID  string  `json:"format_id"`
	Quality   string  `json:"quality"`
	Extension string  `json:"ext"`
	FileSize  int64   `json:"filesize,omitempty"`
	HasVideo  bool    `json:"vcodec"` // Will be parsed from vcodec != "none"
	HasAudio  bool    `json:"acodec"` // Will be parsed from acodec != "none"
	Width     int     `json:"width,omitempty"`
	Height    int     `json:"height,omitempty"`
	FPS       float64 `json:"fps,omitempty"` // Changed from int to float64
}

// FormatSelectionRequest represents a user's format choice
type FormatSelectionRequest struct {
	URL      string
	ChatID   int64
	FormatID string
	Quality  string
}
