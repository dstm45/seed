// Package database gère la connexion à la base de données et les requêtes.
package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

// Connection opens a PostgreSQL database connection using environment variables
func Connection() *pgx.Conn {
	dsn := os.Getenv("DATABASE_URL")

	// Open the database connection
	ctx := context.Background()
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalln("Erreur lors de la connexion à la base de données:", err)
	}
	// Verify the connection by attempting a simple query
	if err := db.Ping(ctx); err != nil {
		log.Fatalln("Erreur lors de la vérification de la connexion à la base de données:", err)
	}

	return db
}
