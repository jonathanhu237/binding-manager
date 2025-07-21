"""create user table

Revision ID: cb909da56976
Revises:
Create Date: 2025-07-20 14:05:59.649413

"""

from typing import Sequence, Union

from alembic import op

# revision identifiers, used by Alembic.
revision: str = "cb909da56976"
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
            hashed_password bytea UNIQUE NOT NULL,
            email citext UNIQUE NOT NULL,
            role text NOT NULL CHECK (role IN ('admin', 'user')) DEFAULT 'user',
            version integer NOT NULL DEFAULT 1
        );
    """)


def downgrade() -> None:
    """Downgrade schema."""
    op.execute("""
        DROP TABLE IF EXISTS users;
               
        DROP EXTENSION IF EXISTS citext;
    """)
