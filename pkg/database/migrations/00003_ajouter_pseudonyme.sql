-- +goose Up
ALTER TABLE users ADD COLUMN pseudonyme varchar(70);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE users DROP COLUMN pseudonyme;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
