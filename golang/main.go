package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Message struct for OpenAI GPT-3
// Role is the role of the message, either "system" or "user"
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIRequest struct for OpenAI GPT-3
// Model is the model name to use
type OpenAIRequest struct {
	Model   string    `json:"model"`
	Message []Message `json:"messages"`
}

// OpenAIResponse struct for OpenAI GPT-3
// Choices is an array of responses from the model
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	// fmt.Println("OpenAI GPT-3 Example")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	organization := os.Getenv("OPENAI_ORGANIZATION")

	fmt.Println("API Key:", apiKey)
	fmt.Println("Organization:", organization)
}
