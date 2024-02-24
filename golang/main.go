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

// Request struct for the OpenAI Chat Completions API
type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message struct for the OpenAI Chat Completions API
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response struct for the HTTP response
type Response struct {
	Choices []Choice `json:"choices"`
}

// Choice struct for the response choices
type Choice struct {
	Message Content `json:"message"`
}

// Content struct for the message content
type Content struct {
	Content string `json:"content"`
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
	} else {
		fmt.Println("\nAPI Key:", apiKey)
	}

	// Infinite loop to continuously prompt the user for input
	for {
		// Prompt the user for input
		fmt.Print("\nEnter your prompt (or 'exit' to quit): ")

		reader := bufio.NewReader(os.Stdin)
		prompt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		prompt = strings.TrimSpace(prompt)

		if strings.ToLower(prompt) == "exit" {
			fmt.Println("\nGoodbye!")
			break
		} else if prompt == "" {
			fmt.Println("\nPlease enter a prompt.")
			continue
		} else {
			fmt.Println("\nUser Prompt:", prompt)
		}

		// Define the request payload
		reqPayload := Request{
			Model: "gpt-3.5-turbo",
			Messages: []Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		// Marshal the request payload into JSON
		reqBody, err := json.Marshal(reqPayload)

		if err != nil {
			log.Fatalf("\nError marshaling request payload: %v \n", err)
		} else {
			fmt.Println("\nRequest Body:", string(reqBody))
		}

		// Create the HTTP request
		req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatalf("\nError creating HTTP request: %v \n", err)
		} else {
			fmt.Println("\nHTTP Request:", req)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		fmt.Println("\nRequest Headers:", req.Header)

		// Send the HTTP request
		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			log.Fatalf("\nError sending HTTP request: %v \n", err)
		} else {
			fmt.Println("\nHTTP Response:", resp)
		}

		fmt.Println("\nResponse Status:", resp.Status)

		// Decode the response body
		var responseBody Response

		err = json.NewDecoder(resp.Body).Decode(&responseBody)

		if err != nil {
			log.Fatalf("\nError decoding response body: %v", err)
		} else {
			fmt.Println("\nResponse Body:", responseBody)
		}

		// Print the content of the first choice
		if len(responseBody.Choices) > 0 {
			fmt.Println("\nChatGPT:", responseBody.Choices[0].Message.Content)
		} else {
			fmt.Println("\nNo response choice received.")
		}

		resp.Body.Close()
	}
}
