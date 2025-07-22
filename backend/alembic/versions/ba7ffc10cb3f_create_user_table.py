"""create user table

Revision ID: ba7ffc10cb3f
Revises: 2e5b5a8634c7
Create Date: 2025-07-21 07:28:13.737875

"""

from typing import Sequence, Union

from alembic import op


# revision identifiers, used by Alembic.
revision: str = "ba7ffc10cb3f"
down_revision: Union[str, Sequence[str], None] = "2e5b5a8634c7"
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
            role_id bigserial NOT NULL references role(id) ON DELETE CASCADE,
            version integer NOT NULL DEFAULT 1
        );
    """)


def downgrade() -> None:
    """Downgrade schema."""
    op.execute("""
        DROP TABLE IF EXISTS users;
               
        DROP EXTENSION IF EXISTS citext;
    """)
