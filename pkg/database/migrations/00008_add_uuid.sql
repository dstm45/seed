-- +goose Up
ALTER TABLE users ADD COLUMN uuid UUID NOT NULL DEFAULT gen_random_uuid();
ALTER TABLE evenements ADD COLUMN uuid UUID NOT NULL DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE users DROP COLUMN uuid;
ALTER TABLE evenements DROP COLUMN uuid;
