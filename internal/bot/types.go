package bot

// User represents a Telegram user or bot
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// GetMeResponse represents the response from getMe API call
type GetMeResponse struct {
	Ok          bool   `json:"ok"`
	Result      User   `json:"result"`
	Description string `json:"description,omitempty"`
}

// Message represents a Telegram message
type Message struct {
	MessageID int64  `json:"message_id"`
	From      *User  `json:"from,omitempty"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text,omitempty"`
}

// Chat represents a Telegram chat
type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// Update represents an incoming update from Telegram
type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

// GetUpdatesResponse represents the response from getUpdates API call
type GetUpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// SendMessageRequest represents a request to send a message
type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessageResponse represents the response from sendMessage API call
type SendMessageResponse struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

// WebhookInfo represents webhook information
type WebhookInfo struct {
	URL                  string `json:"url"`
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int    `json:"pending_update_count"`
	LastErrorDate        int64  `json:"last_error_date,omitempty"`
	LastErrorMessage     string `json:"last_error_message,omitempty"`
}

// GetWebhookInfoResponse represents the response from getWebhookInfo
type GetWebhookInfoResponse struct {
	Ok     bool        `json:"ok"`
	Result WebhookInfo `json:"result"`
}

// SetWebhookRequest represents a request to set webhook
type SetWebhookRequest struct {
	URL string `json:"url"`
}

// SetWebhookResponse represents the response from setWebhook
type SetWebhookResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description,omitempty"`
}
