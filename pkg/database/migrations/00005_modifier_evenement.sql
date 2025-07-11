-- +goose Up
ALTER TABLE evenements ADD COLUMN heure_evenement TIMESTAMP;
ALTER TABLE evenements ADD COLUMN location_evenement TEXT;
ALTER TABLE evenements ADD COLUMN chemin_photo TEXT;
ALTER TABLE evenements ADD COLUMN categorie TEXT;
ALTER TABLE evenements ADD COLUMN prix_billet FLOAT;
ALTER TABLE evenements ADD COLUMN quantite_billet INTEGER;

-- +goose Down
ALTER TABLE evenements DROP COLUMN heure_evenement;
ALTER TABLE evenements DROP COLUMN location_evenement;
ALTER TABLE evenements DROP COLUMN chemin_photo;
ALTER TABLE evenements DROP COLUMN categorie;
ALTER TABLE evenements DROP COLUMN prix_billet;
ALTER TABLE evenements DROP COLUMN quantite_billet;
