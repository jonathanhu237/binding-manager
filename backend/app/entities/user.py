from dataclasses import dataclass


@dataclass
class User:
    id: int
    username: str
    password_hash: str
    email: str
    is_admin: bool
    version: int
