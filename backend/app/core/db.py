from contextlib import contextmanager
from sqlalchemy import create_engine
from app.core.config import settings

engine = create_engine(settings.postgres_database_uri)


@contextmanager
def get_db_connection():
    conn = engine.connect()
    try:
        yield conn
    finally:
        conn.close()
