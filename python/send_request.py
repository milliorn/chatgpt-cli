import requests
import json
from dataclasses import asdict
from typing import Optional

from models import Message, RequestBody


def send_openai_request(api_key: str, user_prompt: str) -> Optional[str]:
    """
    Sends a prompt to the OpenAI chat completion endpoint using the new model (gpt-4)
    and returns the response text, or None if an error occurs.
    """
    # Build the request body using our dataclass, with the new model
    body = RequestBody(
        model="gpt-4",  # Updated model
        messages=[Message(role="user", content=user_prompt)],
    )

    url = "https://api.openai.com/v1/chat/completions"
    headers = {"Content-Type": "application/json", "Authorization": f"Bearer {api_key}"}

    try:
        # Convert our dataclass to JSON and send the request
        response = requests.post(
            url, headers=headers, data=json.dumps(asdict(body)), timeout=10
        )

        # Check if the response status is OK
        if response.status_code != 200:
            print(f"Non-OK status: {response.status_code}")
            print(f"Response Body: {response.text}")
            response.raise_for_status()

        response_json = response.json()

        if "choices" in response_json and len(response_json["choices"]) > 0:
            return response_json["choices"][0]["message"]["content"]
        else:
            return None

    except (requests.RequestException, KeyError) as e:
        print(f"Error during OpenAI request: {e}")
        return None
