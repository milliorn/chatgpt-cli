import os
from dotenv import load_dotenv

from send_request import send_openai_request


def main():
    # Load environment variables from .env
    load_dotenv()

    # Retrieve the OPENAI_API_KEY
    api_key = os.getenv("OPENAI_API_KEY")

    if not api_key:
        print("Error: OPENAI_API_KEY not found in environment.")
        return

    print("Welcome to the Python Chat CLI!")

    while True:
        user_input = input("\nEnter your prompt (or 'exit' to quit): ").strip()

        if user_input.lower() == "exit":
            print("\nGoodbye!")
            break

        if not user_input:
            print("Please enter a prompt.")
            continue

        # print(f"You entered: {user_input}")

        response_text = send_openai_request(api_key, user_input)

        if response_text is not None:
            print(f"\nChatGPT: {response_text}")
        else:
            print("\nNo response received or an error occurred.")


if __name__ == "__main__":
    main()
