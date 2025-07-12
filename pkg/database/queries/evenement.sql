-- name: NewEvent :exec
INSERT INTO 
evenements (nom, description, debut_vente, fin_vente, organisateur, heure_evenement, location_evenement, chemin_photo, categorie, prix_billet, quantite_billet)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: GetEventByUserEmail :many
SELECT ev.* FROM evenements ev
LEFT JOIN users u ON ev.organisateur = u.id
WHERE u.email = $1;

-- name: GetEventByID :one
SELECT * FROM evenements WHERE id=$1;
