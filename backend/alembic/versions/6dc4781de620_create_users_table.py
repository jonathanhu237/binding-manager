"""create users table

Revision ID: 6dc4781de620
Revises:
Create Date: 2025-07-22 16:46:34.129262

"""

from typing import Sequence, Union

from alembic import op

# revision identifiers, used by Alembic.
revision: str = "6dc4781de620"
down_revision: Union[str, Sequence[str], None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    op.execute("""
        CREATE EXTENSION IF NOT EXISTS citext;

        CREATE TABLE users (
            id bigserial PRIMARY KEY,
            username text UNIQUE NOT NULL,
            password_hash bytea UNIQUE NOT NULL,
            email citext UNIQUE NOT NULL,
            is_admin boolean NOT NULL DEFAULT FALSE,
            version integer NOT NULL DEFAULT 1
        );
    """)


def downgrade() -> None:
    """Downgrade schema."""
    op.execute("""
        DROP TABLE IF EXISTS users;
               
        DROP EXTENSION IF EXISTS citext;
    """)
