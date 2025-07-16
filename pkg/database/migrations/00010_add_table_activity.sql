-- +goose Up
CREATE TABLE activites (
  id BIGSERIAL PRIMARY KEY,
  type TEXT NOT NULL,
  responsable_id BIGINT NOT NULL REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP COLUMN activites;
