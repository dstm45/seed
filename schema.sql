CREATE TABLE IF NOT EXISTS agriculteurs(
    id SERIAL PRIMARY KEY,
    nom VARCHAR(45),
    postnom VARCHAR(45),
    prenom VARCHAR(45),
    id_region INTEGER,
    id_stock INTEGER
);

CREATE TABLE IF NOT EXISTS produits(
    id SERIAL PRIMARY KEY,
    nom VARCHAR(45),
    id_agriculteur INTEGER
);

CREATE TABLE IF NOT EXISTS stock(
    id SERIAL PRIMARY KEY,
);

CREATE TABLE IF NOT EXISTS champs(
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS marches(
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50),
    id_region INTEGER
    FOREIGN KEY (id_region) REFERENCES regions(id),
);

CREATE TABLE IF NOT EXISTS commandes(
    id SERIAL PRIMARY KEY,
    id_produit INTEGER,
    id_client INTEGER
);
