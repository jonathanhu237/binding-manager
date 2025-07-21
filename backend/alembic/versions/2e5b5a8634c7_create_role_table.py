"""create role table

Revision ID: 2e5b5a8634c7
Revises:
Create Date: 2025-07-21 07:22:38.506858

"""

from typing import Sequence, Union

from alembic import op


# revision identifiers, used by Alembic.
revision: str = "2e5b5a8634c7"
down_revision: Union[str, Sequence[str], None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    op.execute("""
        CREATE TABLE IF NOT EXISTS role (
            id bigserial PRIMARY KEY,
            name text UNIQUE NOT NULL
        );
    """)


def downgrade() -> None:
    """Downgrade schema."""
    op.execute("""
        DROP TABLE IF EXISTS roles;
    """)
