from dataclasses import dataclass


@dataclass
class User:
    id: int | None = None
    username: str | None = None
    password_hash: str | None = None
    email: str | None = None
    is_admin: bool | None = None
    version: int | None = None
