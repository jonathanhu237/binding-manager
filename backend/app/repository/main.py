from psycopg_pool import ConnectionPool

from app.repository.user import UserRepository


class Repository:
    def __init__(self, pool: ConnectionPool):
        self.user = UserRepository(pool)
