-- name: GetUserByEmail :one
SELECT 
  * 
FROM 
  users 
WHERE 
  email = $1;

-- name: NewUser :one
INSERT INTO 
  users (nom, prenom, email, password_hash, pseudonyme)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateDescription :exec
UPDATE users 
SET 
  nom = $1,
  prenom=$2,
  description=$3,
  chemin_photo=$4,
  pseudonyme=$5
WHERE 
  email=$6;

-- name: UpdatePassword :exec
UPDATE 
  users
SET 
  password_hash = $1
WHERE 
  email=$2;

-- name: GetUserByID :one
SELECT 
  * 
FROM 
  users 
WHERE 
  id=$1;

-- name: GetUserByUUID :one
SELECT
*
FROM 
  users 
WHERE 
  uuid=$1;
