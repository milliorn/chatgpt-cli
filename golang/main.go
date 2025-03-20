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

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env. If we can't, it's a fatal error.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env file:", err)
	}

	// Retrieve the OPENAI_API_KEY from the environment.
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set.")
	}

	// Prepare an HTTP client and a reader for console input.
	client := &http.Client{}
	reader := bufio.NewReader(os.Stdin)

	// Prompt for user input until they type "exit".
	for {
		fmt.Print("\nEnter your prompt (or 'exit' to quit): ")
		
		prompt, err := reader.ReadString('\n')

		if err != nil {
			// More descriptive error message for failed input reading.
			log.Fatal("Could not read your prompt. Please verify console input. Error details:", err)
		}

		prompt = strings.TrimSpace(prompt)

		// If "exit", terminate the loop.
		if strings.ToLower(prompt) == "exit" {
			fmt.Println("\nGoodbye!")
			break
		}
		// If the prompt is empty, prompt again.
		if prompt == "" {
			fmt.Println("Please enter a prompt.")
			continue
		}

		// Build the request payload to send to OpenAI.
		reqPayload := Request{
			Model: "gpt-3.5-turbo",
			Messages: []Message{
				{Role: "user", Content: prompt},
			},
		}

		// Send the request and handle errors.
		responseBody, err := sendOpenAIRequest(client, apiKey, reqPayload)
		if err != nil {
			log.Println("Error from OpenAI request:", err)
			continue
		}

		// Print the first choice if available.
		if len(responseBody.Choices) > 0 {
			fmt.Println("\nChatGPT:", responseBody.Choices[0].Message.Content)
		} else {
			fmt.Println("\nNo response choice received.")
		}
	}
}

/*
sendOpenAIRequest:
1. Marshals the Request to JSON.
2. Creates a POST request (with 10s timeout).
3. Checks for a 200 OK response.
4. Decodes the JSON response into a Response struct.
*/
func sendOpenAIRequest(client *http.Client, apiKey string, payload Request) (*Response, error) {
	// Convert the Request struct to JSON.
	reqBody, err := json.Marshal(payload)

	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	// Create a context with a 10s timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build an HTTP POST request using the context.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(reqBody),
	)

	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	// Set required headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Execute the request.
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}

	defer resp.Body.Close()

	// Ensure the status code is OK.
	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		return nil, fmt.Errorf("non-OK status: %s\nBody: %s",
			resp.Status, buf.String())
	}

	// Decode the response into a Response struct.
	var responseBody Response

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	// Return the parsed response.
	return &responseBody, nil
}
