import os
from dotenv import load_dotenv


def main():
    # Load environment variables from .env
    load_dotenv()

    # Retrieve the OPENAI_API_KEY
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        print("Error: OPENAI_API_KEY not found in environment.")
        return

    print("Success! Found OPENAI_API_KEY.")
    print("Welcome to the Python Chat CLI!")


if __name__ == "__main__":
    main()
