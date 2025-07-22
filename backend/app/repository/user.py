from psycopg_pool import ConnectionPool

from app.domain.user import User
from app.repository.error import UnknownError


class UserRepository:
    def __init__(self, pool: ConnectionPool):
        self.pool = pool

    def is_admin_exists(self) -> bool:
        with self.pool.connection() as conn:
            record = conn.execute(
                "SELECT EXISTS(SELECT 1 FROM users WHERE is_admin = True);"
            ).fetchone()

            if record is None:
                raise UnknownError("Failed to check if admin exists in the database.")

            return record[0]

    def create_user(self, user: User):
        with self.pool.connection() as conn:
            record = conn.execute(
                """
                INSERT INTO users (username, password_hash, email, is_admin)
                VALUES (%s, %s, %s, %s)
                RETURNING id, version;
            """,
                (user.username, user.password_hash, user.email, user.is_admin),
            ).fetchone()

            if record is None:
                raise UnknownError("Failed to create user in the database.")

            user.id = record[0]
            user.version = record[1]

            return user
