from app.core.db import get_db_connection
from sqlalchemy import text


def admin_user_exists() -> bool:
    with get_db_connection() as conn:
        result = conn.execute(
            text("""
                SELECT EXISTS(SELECT 1
                    FROM users u
                            JOIN role r ON u.role_id = r.id
                    WHERE r.name = 'admin');
            """)
        )
        # TODO: print the result
        return result.scalar() or False
