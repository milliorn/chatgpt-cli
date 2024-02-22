package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	fmt.Println("OpenAI GPT-3 Chatbot in Go!")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	organization := os.Getenv("OPENAI_ORGANIZATION")

	fmt.Println("API Key:", apiKey)
	fmt.Println("Organization:", organization)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a prompt: ")

	for {
		fmt.Printf("User: ")
		// Read the user's input
		input, _ := reader.ReadString('\n')

		// Create a new OpenAIRequest
		request := OpenAIRequest{
			Model: "gpt-3.5-turbo",
			Message: []Message{
				{
					Role:    "user",
					Content: input,
				},
			},
		}

		// Marshal the request into JSON
		requestBytes, err := json.Marshal(request)

		if err != nil {
			log.Fatal(err)
		}

		// Create a new HTTP request
		req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(requestBytes))

		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		// Set the request headers
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("OpenAI-Organization", organization)

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request to OpenAI: %v", err)
		}

		// Close the response body when the function returns
		// defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Decode the response JSON into the OpenAIResponse struct
		var response OpenAIResponse

		// Unmarshal the response JSON into the OpenAIResponse struct
		err = json.Unmarshal(body, &response)

		if err != nil {
			log.Fatalf("Error unmarshalling response: %v", err)
		}

		// Print the response
		fmt.Println("Response:", response.Choices[0].Message.Content)

		if len(response.Choices) > 0 && response.Choices[0].Message.Content != "" {
			fmt.Printf("System: %s", response.Choices[0].Message.Content)
		} else {
			fmt.Println("System: I'm sorry, I don't understand.")
		}
		// // Prompt the user for input
		// fmt.Printf("Enter a prompt: ")
	}
}
