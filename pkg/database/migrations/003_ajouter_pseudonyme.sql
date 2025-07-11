-- +goose Up
ALTER TABLE users ADD COLUMN pseudonyme varchar(70);

-- +goose Down
ALTER TABLE users DROP COLUMN pseudonyme;
