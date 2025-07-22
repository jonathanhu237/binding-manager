postgres_dsn := "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}"

new-migration name:
    migrate create -ext sql -dir migrations -seq {{name}}

migrate-up step:
    migrate -database {{postgres_dsn}} -path migrations up {{step}}

migrate-down step:
    migrate -database {{postgres_dsn}} -path migrations down {{step}}

migration-version:
    migrate -database {{postgres_dsn}} -path migrations version

set-migration-version version:
    migrate -database {{postgres_dsn}} -path migrations force {{version}}