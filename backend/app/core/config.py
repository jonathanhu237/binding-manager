from pydantic_settings import BaseSettings, SettingsConfigDict


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

    model_config = SettingsConfigDict()


def get_config() -> Settings:
    return Settings.model_validate({})


if __name__ == "__main__":
    print(get_config().model_dump_json(indent=4))
