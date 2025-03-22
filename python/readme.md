# Python Chat CLI with OpenAI GPT-4

This is a Python-based command-line interface (CLI) for interacting with OpenAI's ChatGPT (GPT-4) model. It provides a straightforward way to communicate with GPT-4 directly from your terminal, making it ideal for testing, quick prototyping, or just exploring conversational AI.

## Features

- **Interactive Conversations:** Engage in natural, real-time dialogues with GPT-4.
- **Easy Setup:** Quickly configure using a `.env` file.
- **Simple Interface:** Minimal commands required—just type and chat.
- **Robust Error Handling:** Clearly informs you of configuration and API issues.

## Requirements

- **Python 3.7+**
- An active OpenAI account and API key ([sign up here](https://platform.openai.com/)).

## Installation

Clone this repository and install dependencies:

```bash
pip install -r requirements.txt
```

Create a .env file in the project root directory and add your API key:

```ini
OPENAI_API_KEY=your_openai_api_key_here
```

## Usage

Start chatting with GPT-4 from your terminal:

```sh
python chat_cli.py
```

Enter your prompt and ChatGPT will respond. To end the conversation, type:

```sh
exit
```

## Acknowledgements

OpenAI for access to the GPT-4 model.

python-dotenv for handling environment variables.

Requests for HTTP requests to OpenAI API.
