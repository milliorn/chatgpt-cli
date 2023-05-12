# chatgpt-cli

[![CodeQL](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/milliorn/chatgpt-cli/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/milliorn/chatgpt-cli/actions/workflows/dependency-review.yml)

A command-line interface (CLI) tool that utilizes OpenAI's GPT-3 language model for interactive chatbot conversations.
With this tool, you can quickly generate natural language responses to prompts from your terminal using the powerful GPT-3 API.

## Usage

To start a conversation with the GPT-3 chatbot, simply type the following command in your terminal in project root:

`npm start`

or

`node chatgpt`

The chatbot will generate a prompt for you to respond to, and it will continue the conversation with you based on your input.

## Configuration

Before you can use the `chatgpt-cli` tool, you'll need to set up an OpenAI API key. 
To do this, create an account on the OpenAI website and follow the instructions for obtaining an API key.
Then, you need to create an `.env` file in project root, an add your organization and key there. 

`API_KEY=yourkey`

`ORGANIZATION=org`
