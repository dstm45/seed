-- name: NewEvent :exec
INSERT INTO 
  evenements (nom, description, debut_vente, fin_vente, organisateur, heure_evenement, location_evenement, chemin_photo, categorie, prix_billet, quantite_billet)
VALUES 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: GetEventByUserEmail :many
SELECT 
  ev.* 
FROM 
  evenements ev
LEFT JOIN 
  users u ON ev.organisateur = u.id
WHERE 
  u.email = $1;

-- name: GetEventByUUID :one
SELECT
  * 
FROM 
  evenements 
WHERE 
  uuid=$1;

-- name: GetAllEvent :many
SELECT 
  * 
FROM 
  evenements;

-- name: UpdateEventByUUID :exec
UPDATE evenements
SET
  nom = $1,
  description = $2,
  debut_vente = $3,
  fin_vente = $4,
  organisateur = $5,
  heure_evenement = $6,
  location_evenement = $7,
  chemin_photo = $8,
  categorie = $9,
  prix_billet = $10,
  quantite_billet = $11
WHERE uuid = $12;

-- name: DeleteEventByUUID :exec
DELETE FROM 
  evenements 
WHERE 
  uuid=$1;

