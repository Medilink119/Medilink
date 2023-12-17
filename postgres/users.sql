CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password_hash bytea NOT NULL
);