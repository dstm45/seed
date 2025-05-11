-- name: GetUserByEmail :one
SELECT * from users WHERE email = ?;
-- name: GetUserId :one
SELECT * from users WHERE id = ?;
-- name: GetPasswordHash :one
SELECT password_hash from users WHERE email = ?;

-- name: NewUser :exec
INSERT INTO users (nom, postnom, prenom, email, password_hash)VALUES(?, ?, ?, ?, ?);