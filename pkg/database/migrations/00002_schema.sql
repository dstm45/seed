-- +goose Up
ALTER TABLE users ALTER COLUMN password_hash TYPE TEXT;
ALTER TABLE users DROP COLUMN postnom;

-- +goose Down
ALTER TABLE users ALTER COLUMN password_hash TYPE VARCHAR(200);
ALTER TABLE users ADD COLUMN postnom VARCHAR(255);