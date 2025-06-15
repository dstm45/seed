-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(45),
    postnom VARCHAR(45),
    prenom VARCHAR(45),
    email VARCHAR(100),
    type_compte VARCHAR(20),
    password_hash VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS evenements (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(45),
    description TEXT NOT NULL,
    debut_vente DATE NOT NULL,
    fin_vente DATE NOT NULL,
    date_evenement DATE NOT NULL,
    organisateur BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (organisateur) REFERENCES users(id)
);

-- +goose Down
DROP TABLE evenements;
DROP TABLE users;