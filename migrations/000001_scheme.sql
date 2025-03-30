-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id       SERIAL       PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    name     VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    age      INTEGER      NOT NULL,
    verified BOOLEAN      DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
