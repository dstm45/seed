-- name: GetUserByEmail :one
SELECT * from users WHERE email = ?;
-- name: GetUserById :one
SELECT * from users WHERE id = ?;
-- name: GetPasswordHash :one
SELECT password_hash from users WHERE email = ?;

-- name: GetUsers :many
SELECT * from users;

-- name: NewUser :exec
INSERT INTO users (nom, postnom, prenom, email, password_hash)
VALUES(?, ?, ?, ?, ?);