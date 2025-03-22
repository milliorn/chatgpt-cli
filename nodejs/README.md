# chatgpt-cli

[![CodeQL](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql)

ChatGPT CLI is a command-line interface (CLI) tool that allows you to have interactive conversations with OpenAI's GPT-3 powered chatbot, known as ChatGPT. With this CLI tool, you can have natural language conversations with the chatbot and receive text-based responses.

The application is built using Python and leverages OpenAI's GPT-3 API to enable seamless communication with the chatbot. It provides a simple and efficient way to interact with the language model, making it suitable for testing, prototyping, and exploring various use cases.

## Features

- Interactive Chat: Engage in real-time, back-and-forth conversations with ChatGPT.
- Easy-to-Use: Simple and intuitive command-line interface for smooth interactions.
- Customizable Prompts: Tailor your conversation starters and prompts as per your requirements.

## Usage

1. Install the required dependencies: `npm install`
2. To start a conversation with the GPT-3 chatbot, simply type the following command in your terminal in project root:

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

## Acknowledgements

This project was built using the [dotenv](https://www.npmjs.com/package/dotenv) and [openai](https://www.npmjs.com/package/openai) NPM packages. Special thanks to the contributors of these packages for making this project possible.

Feel free to contribute to the project by opening issues or submitting pull requests.

Happy chatting with ChatGPT! 🚀
