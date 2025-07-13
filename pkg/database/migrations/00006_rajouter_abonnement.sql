-- +goose Up
CREATE TABLE abonnements (
  id BIGSERIAL PRIMARY KEY,
  creatorid BIGINT NOT NULL,
  followerid BIGINT NOT NULL,
  CONSTRAINT fk_creator FOREIGN KEY (creatorid) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_follower FOREIGN KEY (followerid) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT unique_subscription UNIQUE (creatorid, followerid)
);

-- +goose Down
DROP TABLE abonnements;
