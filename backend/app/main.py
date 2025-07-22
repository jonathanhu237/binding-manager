from contextlib import asynccontextmanager
import logging
from fastapi import FastAPI
from psycopg_pool import AsyncConnectionPool

from app.api.main import api_router
from app.config import get_config

# Get the logger of FastAPI
logger = logging.getLogger("uvicorn.error")

# Load the configuration
config = get_config()

# Define the connection pool
pool = AsyncConnectionPool(config.postgres_dsn, open=False)


# Initialize the connection pool
@asynccontextmanager
async def lifespan(instance: FastAPI):
    await pool.open()
    logger.info("Database connection pool initialized.")
    yield
    await pool.close()


app = FastAPI(title="Binding Manager API", lifespan=lifespan)
app.include_router(api_router)
