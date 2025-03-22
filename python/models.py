from dataclasses import dataclass
from typing import List


@dataclass
class Message:
    role: str
    content: str


@dataclass
class RequestBody:
    model: str
    messages: List[Message]
