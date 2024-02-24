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
	Model    string    `json:"model,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

// Message struct for the OpenAI Chat Completions API
type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// Response struct for the HTTP response
type Response struct {
	Choices []Choice `json:"choices,omitempty"`
}

// Choice struct for the response choices
type Choice struct {
	Message Content `json:"message,omitempty"`
}

// Content struct for the message content
type Content struct {
	Content string `json:"content,omitempty"`
}

func main() {
	fmt.Println("OpenAI GPT-3 Chatbot in Go!")

	// Load environment variables
	err := godotenv.Load()

	handleEnvLoadError(err)

	// Load API key from environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")

	checkAPIKey(apiKey)

	// Infinite loop to continuously prompt the user for input
	for {
		// Prompt the user for input
		fmt.Print("\nEnter your prompt (or 'exit' to quit): ")

		reader := bufio.NewReader(os.Stdin)

		prompt, err := reader.ReadString('\n')

		checkInputError(err)

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
		reqPayload := createRequestPayload(prompt)

		// Marshal the request payload into JSON
		reqBody, err := json.Marshal(reqPayload)

		handleRequestBody(err, reqBody)

		// Create the HTTP request
		req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))

		handleHTTPRequestCreationError(err, req)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		// fmt.Println("\nRequest Headers:", req.Header)

		// Send the HTTP request
		client := &http.Client{}

		resp, err := client.Do(req)

		handleHTTPResponseSendingError(err, resp)

		// fmt.Println("\nResponse Status:", resp.Status)

		// Decode the response body
		var responseBody Response

		err = json.NewDecoder(resp.Body).Decode(&responseBody)

		handleResponseBodyDecodingError(err, responseBody)

		// Print the content of the first choice
		printResponseChoice(responseBody)

		resp.Body.Close()
	}
}

// handleEnvLoadError is a function that handles errors when loading the .env file.
// It takes an error as a parameter and logs a fatal error message if the error is not nil.
func handleEnvLoadError(err error) {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// checkAPIKey checks if the OpenAI API key is set.
// If the API key is empty, it logs a fatal error message.
// Otherwise, it prints the API key.
func checkAPIKey(apiKey string) {
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set.")
	}
	// else {
	// 	fmt.Println("\nAPI Key:", apiKey)
	// }
}

// checkInputError checks if there is an error while reading input.
// If there is an error, it logs a fatal error message.
func checkInputError(err error) {
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}

// createRequestPayload creates a new Request payload with the provided prompt.
func createRequestPayload(prompt string) Request {
	reqPayload := Request{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	return reqPayload
}

// handleRequestBody handles the request body based on the error status.
// If there's an error, it logs a fatal error message; otherwise, it prints the request body.
func handleRequestBody(err error, reqBody []byte) {
	if err != nil {
		log.Fatalf("\nError marshaling request payload: %v \n", err)
	}
	// else {
	// 	fmt.Println("\nRequest Body:", string(reqBody))
	// }
}

// handleHTTPRequestCreationError handles the error when creating an HTTP request.
// If there's an error, it logs a fatal error message; otherwise, it prints the HTTP request.
func handleHTTPRequestCreationError(err error, req *http.Request) {
	if err != nil {
		log.Fatalf("\nError creating HTTP request: %v \n", err)
	}
	// else {
	// 	fmt.Println("\nHTTP Request:", req)
	// }
}

// handleHTTPResponseSendingError handles the error when sending an HTTP request.
// If there's an error, it logs a fatal error message; otherwise, it prints the HTTP response.
func handleHTTPResponseSendingError(err error, resp *http.Response) {
	if err != nil {
		log.Fatalf("\nError sending HTTP request: %v \n", err)
	}
	// else {
	// 	fmt.Println("\nHTTP Response:", resp)
	// }
}

// handleResponseBodyDecodingError handles the error when decoding the response body.
// If there's an error, it logs a fatal error message; otherwise, it prints the response body.
func handleResponseBodyDecodingError(err error, responseBody Response) {
	if err != nil {
		log.Fatalf("\nError decoding response body: %v", err)
	}
	// else {
	// 	fmt.Println("\nResponse Body:", responseBody)
	// }
}

// printResponseChoice prints the content of the first response choice if available.
// If no response choice is received, it prints a message indicating that.
func printResponseChoice(responseBody Response) {
	if len(responseBody.Choices) > 0 {
		fmt.Println("\nChatGPT:", responseBody.Choices[0].Message.Content)
	} else {
		fmt.Println("\nNo response choice received.")
	}
}
