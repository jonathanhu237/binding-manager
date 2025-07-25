set dotenv-load := true
set dotenv-required := true

postgres_dsn := "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

[group('database migration')]
new-migration name:
    docker compose run --rm --user "$(id -u):$(id -g)" migrator create -ext sql -dir migrations -seq {{name}}

[group('database migration')]
migrate-up step:
    docker compose run --rm migrator -database {{postgres_dsn}} -path migrations up {{step}}

[group('database migration')]
migrate-down step:
    docker compose run --rm migrator -database {{postgres_dsn}} -path migrations down {{step}}

[group('database migration')]
migration-version:
    docker compose run --rm migrator -database {{postgres_dsn}} -path migrations version

[group('database migration')]
set-migration-version version:
    docker compose run --rm migrator -database {{postgres_dsn}} -path migrations force {{version}}
