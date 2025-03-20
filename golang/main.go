package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type (
	// Request for the OpenAI Chat Completions API
	Request struct {
		Model    string    `json:"model,omitempty"`
		Messages []Message `json:"messages,omitempty"`
	}

	// Message for the OpenAI Chat Completions API
	Message struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	}

	// Response from the OpenAI Chat Completions API
	Response struct {
		Choices []Choice `json:"choices,omitempty"`
	}

	// Choice within the response
	Choice struct {
		Message Content `json:"message,omitempty"`
	}

	// Content struct for the message content
	Content struct {
		Content string `json:"content,omitempty"`
	}
)

func main() {
	fmt.Println("OpenAI GPT-3 Chatbot in Go!")

	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// 2. Check for the required OPENAI_API_KEY
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set.")
	}

	client := &http.Client{}
	reader := bufio.NewReader(os.Stdin)

	for {
		// 3. Read user input
		fmt.Print("\nEnter your prompt (or 'exit' to quit): ")
		prompt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input: ", err)
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

		// 4. Build the request payload
		reqPayload := Request{
			Model: "gpt-3.5-turbo",
			Messages: []Message{
				{Role: "user", Content: prompt},
			},
		}

		// 5. Send the request to OpenAI
		responseBody, err := sendOpenAIRequest(client, apiKey, reqPayload)
		if err != nil {
			log.Println("Error from OpenAI request:", err)
			continue
		}

		// 6. Print the response
		if len(responseBody.Choices) > 0 {
			fmt.Println("\nChatGPT:", responseBody.Choices[0].Message.Content)
		} else {
			fmt.Println("\nNo response choice received.")
		}
	}
}

// sendOpenAIRequest marshals the request payload and sends an HTTP POST to the OpenAI API.
func sendOpenAIRequest(client *http.Client, apiKey string, payload Request) (*Response, error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Optionally, read the body for more info
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
