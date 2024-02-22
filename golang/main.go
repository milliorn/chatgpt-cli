package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Request struct for the OpenAI Chat Completions API
type Request struct {
	Model          string     `json:"model"`
	ResponseFormat FormatType `json:"response_format"`
	Messages       []Message  `json:"messages"`
}

// Message struct for the OpenAI Chat Completions API
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// FormatType struct for the OpenAI Chat Completions API
type FormatType struct {
	Type string `json:"type"`
}

func main() {
	fmt.Println("OpenAI GPT-3 Chatbot in Go!")

	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set.")
	}

	fmt.Println("API Key:", apiKey)

	// Define the request payload
	reqPayload := Request{
		Model: "gpt-3.5-turbo",
		ResponseFormat: FormatType{
			Type: "json_object",
		},
		Messages: []Message{
			{
				Role:    "user",
				Content: "Hello, how are you?",
			},
			{
				Role:    "system",
				Content: "This is a message containing the word 'json'.",
			},
		},
	}

	fmt.Println("Request Payload:", reqPayload)

	// Marshal the request payload into JSON
	reqBody, err := json.Marshal(reqPayload)

	if err != nil {
		log.Fatalf("Error marshaling request payload: %v \n", err)
	} else {
		// Print the request body
		fmt.Println("Request Body:", reqBody)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatalf("Error creating HTTP request: %v + \n", err)
	} else {
		// Print the HTTP request
		fmt.Println("Request:", req)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	fmt.Println("Request Headers:", req.Header)

	// Send the HTTP request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error sending HTTP request: %v \n", err)
	} else {
		// Print the HTTP response
		fmt.Println("Response:", resp)
	}

	defer resp.Body.Close()

	// Read the response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	} else {
		// Print the response body
		fmt.Println("Response Body:", result)
	}

	// Check if the "choices" field exists and contains elements
	choices, ok := result["choices"].([]interface{})

	if !ok || len(choices) == 0 {
		log.Fatal("No response received from OpenAI.")
	} else {
		// Print the response choices
		fmt.Println("Response Choices:", choices)
	}

	// Extract and print the response content
	content, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	if !ok {
		log.Fatal("Error parsing response content")
	}
	fmt.Println("Response:", content)
}
