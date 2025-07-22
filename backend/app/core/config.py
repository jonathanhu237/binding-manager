from pydantic import PostgresDsn, computed_field
from pydantic_settings import BaseSettings, SettingsConfigDict
import urllib.parse


class Settings(BaseSettings):
    jwt_secret: str
    jwt_expire_minutes: int

    first_admin_username: str
    first_admin_password: str
    first_admin_email: str

    postgres_username: str
    postgres_password: str
    postgres_host: str
    postgres_port: int
    postgres_db: str

    @computed_field
    @property
    def postgres_dsn(self) -> str:
        return str(
            PostgresDsn.build(
                scheme="postgresql",
                username=self.postgres_username,
                password=urllib.parse.quote_plus(self.postgres_password),
                host=self.postgres_host,
                port=self.postgres_port,
                path=self.postgres_db,
            )
        )

    model_config = SettingsConfigDict()


if __name__ == "__main__":
    print(Settings.model_validate({}).model_dump_json(indent=4))
