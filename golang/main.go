package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	// Loads environment variables from a .env file if present.
	"github.com/joho/godotenv"
)

type (
	/*
	   Request contains the model name and a list of user messages
	   to send to the OpenAI "chat/completions" endpoint.
	*/
	Request struct {
		Model    string    `json:"model,omitempty"`
		Messages []Message `json:"messages,omitempty"`
	}

	/*
	   Message represents a single conversation entry with a specific role
	   (user, assistant, or system) and its text content.
	*/
	Message struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	}

	/*
	   Response is what we get back from the OpenAI API and contains a list
	   of possible reply choices.
	*/
	Response struct {
		Choices []Choice `json:"choices,omitempty"`
	}

	/*
	   Choice is one possible reply from the API, which includes the final
	   generated text in the Message field.
	*/
	Choice struct {
		Message Content `json:"message,omitempty"`
	}

	/*
	   Content holds the actual text returned by the OpenAI model.
	*/
	Content struct {
		Content string `json:"content,omitempty"`
	}
)

func main() {
	fmt.Println("OpenAI GPT-3 Chatbot in Go!")

	// Load environment variables (e.g., OPENAI_API_KEY) from .env.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get the API key from the environment.
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set.")
	}

	// Prepare an HTTP client and a reader for console input.
	client := &http.Client{}
	reader := bufio.NewReader(os.Stdin)

	// Continuously prompt for user input until "exit" is typed.
	for {
		fmt.Print("\nEnter your prompt (or 'exit' to quit): ")
		prompt, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Error reading input:", err)
		}

		prompt = strings.TrimSpace(prompt)

		if strings.ToLower(prompt) == "exit" {
			fmt.Println("\nGoodbye!")
			break
		}

		if prompt == "" {
			fmt.Println("Please enter a prompt.")
			continue
		}

		// Create a request payload with the chosen model and user's prompt.
		reqPayload := Request{
			Model: "gpt-3.5-turbo",
			Messages: []Message{
				{Role: "user", Content: prompt},
			},
		}

		// Send request to OpenAI and handle any errors.
		responseBody, err := sendOpenAIRequest(client, apiKey, reqPayload)

		if err != nil {
			log.Println("Error from OpenAI request:", err)
			continue
		}

		// Print the model's first response if available.
		if len(responseBody.Choices) > 0 {
			fmt.Println("\nChatGPT:", responseBody.Choices[0].Message.Content)
		} else {
			fmt.Println("\nNo response choice received.")
		}
	}
}

/*
sendOpenAIRequest handles:
- Converting a Request to JSON,
- Creating a context with a 10s timeout,
- Making a POST to OpenAI using http.NewRequestWithContext,
- Checking the response status code,
- Decoding the response into a Go struct.
*/
func sendOpenAIRequest(client *http.Client, apiKey string, payload Request) (*Response, error) {
	// Convert payload to JSON
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	// Create a context that will auto-cancel after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build the HTTP request with context
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		// If the error is because of the context timing out or being canceled,
		// err may be `context.DeadlineExceeded` or `context.Canceled`
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return nil, fmt.Errorf("non-OK HTTP status: %s\nResponse body: %s",
			resp.Status, buf.String())
	}

	var responseBody Response
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return &responseBody, nil
}
