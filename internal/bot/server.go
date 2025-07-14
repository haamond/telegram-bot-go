package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Server represents the webhook server
type Server struct {
	client *Client
	port   string
}

// NewServer creates a new webhook server
func NewServer(client *Client, port string) *Server {
	return &Server{
		client: client,
		port:   port,
	}
}

// Start starts the webhook server
func (s *Server) Start() error {
	http.HandleFunc("/webhook", s.webhookHandler)
	http.HandleFunc("/health", s.healthHandler)

	fmt.Printf("Starting webhook server on port %s\n", s.port)
	fmt.Printf("Webhook endpoint: /webhook\n")
	fmt.Printf("Health check: /health\n")

	return http.ListenAndServe(":"+s.port, nil)
}

// webhookHandler handles incoming webhook requests from Telegram
func (s *Server) webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var update Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Handle the message if it exists
	if update.Message != nil {
		fmt.Printf("Received webhook message from %s: %s\n",
			update.Message.From.FirstName,
			update.Message.Text)

		err := s.client.HandleMessage(update.Message)
		if err != nil {
			log.Printf("Error handling message: %v", err)
			// Don't return error to Telegram, just log it
		}
	}

	// Always respond with 200 OK to Telegram
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// healthHandler provides a health check endpoint
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Bot is healthy!"))
	if err != nil {
		log.Printf("Error writing health response: %v", err)
	}

}
