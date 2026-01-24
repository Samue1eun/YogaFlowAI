package ai

import (
	"os"

	"github.com/liushuangls/go-anthropic/v2"
)

var Client *anthropic.Client

// InitClient initializes the Anthropic Claude client
func InitClient() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		panic("ANTHROPIC_API_KEY environment variable is required")
	}
	Client = anthropic.NewClient(apiKey)
}

// GetClient returns the initialized Claude client
func GetClient() *anthropic.Client {
	if Client == nil {
		InitClient()
	}
	return Client
}
