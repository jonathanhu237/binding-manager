from logging import Logger
from app.core.config import Settings
from app.repository.main import Repository
from app.service.user import UserService


class Service:
    def __init__(self, logger: Logger, config: Settings, repository: Repository):
        self.user = UserService(logger, config, repository)
