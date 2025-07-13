-- +goose Up
ALTER TABLE evenements ADD COLUMN likes INTEGER NOT NULL DEFAULT 0;
CREATE TABLE commentaires (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  evenement_id BIGINT NOT NULL,
  comment TEXT NOT NULL,
  CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_comment_evenement FOREIGN KEY (evenement_id) REFERENCES evenements(id) ON DELETE CASCADE
);

-- +goose Down
ALTER TABLE evenements DROP COLUMN like;
DROP TABLE commentaires;
