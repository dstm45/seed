-- name: GetUserByEmail :one
SELECT * from users WHERE email = $1;

-- name: GetUserById :one
SELECT * from users WHERE id = $1;

-- name: GetPasswordHash :one
SELECT password_hash from users WHERE email = $1;

-- name: GetUsers :many
SELECT * from users;

-- name: NewUser :exec
INSERT INTO users (nom, prenom, email, password_hash)
VALUES($1, $2, $3, $4);