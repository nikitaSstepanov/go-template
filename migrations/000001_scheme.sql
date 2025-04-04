-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM (
            'USER', 
            'ADMIN'
        );
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id       SERIAL       PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    name     VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    age      INTEGER      NOT NULL,
    role     role         DEFAULT 'USER',
    verified BOOLEAN      DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS role;
-- +goose StatementEnd
