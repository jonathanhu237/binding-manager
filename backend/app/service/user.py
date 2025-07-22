from logging import Logger

import bcrypt
from app.repository.main import Repository
from app.core.config import Settings
from app.domain.user import User


class UserService:
    def __init__(
        self, logger: Logger, config: Settings, repository: Repository
    ) -> None:
        self.logger = logger
        self.config = config
        self.repository = repository

    def ensure_admin_exists(self):
        if self.repository.user.is_admin_exists() is False:
            self.logger.warning("No admin user found. Creating one.")

            admin_user = User(
                username=self.config.first_admin_username,
                password_hash=str(
                    bcrypt.hashpw(
                        self.config.first_admin_password.encode(), bcrypt.gensalt()
                    )
                ),
                email=self.config.first_admin_email,
                is_admin=True,
            )
            self.repository.user.create_user(admin_user)
            self.logger.info("Admin user created successfully.")
