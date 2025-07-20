from pydantic import PostgresDsn, computed_field
from pydantic_core import MultiHostUrl
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    jwt_secret: str
    jwt_expire_minutes: int = 7 * 24 * 60

    first_superuser_username: str = "admin"
    first_superuser_password: str
    first_superuser_email: str

    postgres_server: str = "localhost"
    postgres_port: int = 5432
    postgres_user: str = "binding_manager"
    postgres_password: str
    postgres_db: str = "binding_manager"

    @computed_field
    @property
    def postgres_database_uri(self) -> PostgresDsn:
        return MultiHostUrl.build(
            scheme="postgresql+psycopg2",
            username=self.postgres_user,
            password=self.postgres_password,
            host=self.postgres_server,
            port=self.postgres_port,
            path=self.postgres_db,
        )  # type: ignore

    model_config = SettingsConfigDict()


settings = Settings()  # type: ignore
