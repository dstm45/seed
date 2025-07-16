-- +goose Up
ALTER TABLE users ADD COLUMN online_status BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users DROP COLUMN online_status;
