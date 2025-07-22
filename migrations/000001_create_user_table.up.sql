CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username text UNIQUE NOT NULL,
    password_hash bytea UNIQUE NOT NULL,
    email citext UNIQUE NOT NULL,
    is_admin boolean NOT NULL DEFAULT FALSE,
    version integer NOT NULL DEFAULT 1
);