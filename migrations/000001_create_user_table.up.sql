CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username text UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    email citext UNIQUE NOT NULL,
    is_admin boolean NOT NULL DEFAULT FALSE,
    version integer NOT NULL DEFAULT 1
);