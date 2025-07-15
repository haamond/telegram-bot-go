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
