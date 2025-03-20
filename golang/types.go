package main

// Request holds the model name and a list of user messages for the OpenAI endpoint.
type Request struct {
	Model    string    `json:"model,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

// Message represents a conversation message with a role and its text content.
type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// Response is what we get from the OpenAI API: a list of possible reply choices.
type Response struct {
	Choices []Choice `json:"choices,omitempty"`
}

// Choice contains the actual text returned by the OpenAI model in the Message field.
type Choice struct {
	Message Content `json:"message,omitempty"`
}

// Content represents the text output from the OpenAI model.
type Content struct {
	Content string `json:"content,omitempty"`
}
