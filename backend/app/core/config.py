from pydantic import computed_field
from pydantic_settings import BaseSettings, SettingsConfigDict
from sqlalchemy import URL


class Settings(BaseSettings):
    jwt_secret: str
    jwt_expire_minutes: int = 7 * 24 * 60

    first_admin_username: str = "admin"
    first_admin_password: str
    first_admin_email: str

    postgres_username: str = "binding_manager"
    postgres_password: str
    postgres_host: str = "localhost"
    postgres_port: int = 5432
    postgres_db: str = "binding_manager"

    @computed_field
    @property
    def postgres_database_uri(self) -> str:
        url = URL.create(
            "postgresql+psycopg2",
            username=self.postgres_username,
            password=self.postgres_password,
            host=self.postgres_host,
            port=self.postgres_port,
            database=self.postgres_db,
        ).render_as_string(hide_password=False)

        return str(url)

    model_config = SettingsConfigDict()


settings = Settings()  # type: ignore
