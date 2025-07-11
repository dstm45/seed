-- +goose Up
ALTER TABLE users ADD COLUMN description TEXT;
ALTER TABLE users ADD COLUMN chemin_photo TEXT;

-- +goose Down
ALTER TABLE users DROP COLUMN description;
ALTER TABLE users DROP COLUMN chemin_photo;
