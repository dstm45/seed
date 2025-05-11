-- +goose Up
ALTER TABLE users
ADD COLUMN email VARCHAR(255) UNIQUE NOT NULL,
ADD COLUMN password_hash VARCHAR(255);

-- +goose Down
ALTER TABLE users
DROP COLUMN email;