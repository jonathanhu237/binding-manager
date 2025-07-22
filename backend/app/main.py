import logging
from contextlib import asynccontextmanager

from fastapi import FastAPI
from psycopg import Connection
from psycopg.rows import TupleRow
from psycopg_pool import ConnectionPool

from app.core.config import Settings
from app.repository.main import Repository
from app.service.main import Service


@asynccontextmanager
async def lifespan(_: FastAPI):
    # Get the logger of FastAPI
    logger = logging.getLogger("uvicorn.error")

    # Load the configuration
    config = Settings.model_validate({})

    # Initialize the connection pool
    pool: ConnectionPool[Connection[TupleRow]] = ConnectionPool(
        config.postgres_dsn, open=True
    )
    logger.info("Database connection pool initialized.")

    # Create the Repository and Service instances
    repository = Repository(pool)
    service = Service(logger, config, repository)

    # Ensure the admin exists
    service.user.ensure_admin_exists()

    yield

    # Cleanup: close the connection pool
    pool.close()


app = FastAPI(title="Binding Manager API", lifespan=lifespan)
