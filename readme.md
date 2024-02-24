# OpenAI GPT-3 Chatbot

[![CodeQL](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/milliorn/chatgpt-cli/actions/workflows/github-code-scanning/codeql)

This repository contains implementations of a chatbot powered by OpenAI's GPT-3.5 model in different programming languages. Each implementation provides a command-line interface for users to interact with the chatbot in a conversational manner.

ChatGPT CLI is a command-line interface (CLI) tool that allows you to have interactive conversations with OpenAI's GPT-3 powered chatbot, known as ChatGPT. With this CLI tool, you can have natural language conversations with the chatbot and receive text-based responses.

## Features

- Interactive Chat: Engage in real-time, back-and-forth conversations with ChatGPT.
- Easy-to-Use: Simple and intuitive command-line interface for smooth interactions.
- Customizable Prompts: Tailor your conversation starters and prompts as per your requirements.

## Configuration

Before you can use the `chatgpt-cli` tool, you'll need to set up an OpenAI API key.
To do this, create an account on the OpenAI website and follow the instructions for obtaining an API key.
Then, you need to create an `.env` file in the directory of the language you wish to use, an add your organization and key there.

`API_KEY=yourkey`

`ORGANIZATION=org`

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgements

- [dotenv](https://www.npmjs.com/package/dotenv): For managing environment variables.
- [OpenAI](https://openai.com/) - For providing access to the GPT-3.5 model.
- [Joho/godotenv](https://github.com/joho/godotenv) - For the Go library for reading environment variables from a .env file.
- [Golang](https://golang.org/) - For the Go programming language.

Feel free to contribute to the project by opening issues or submitting pull requests.

Happy chatting with ChatGPT! 🚀
