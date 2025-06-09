package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Connection opens a MySQL database connection using environment variables
func Connection() *Queries {
	dsn := os.Getenv("DATABASE_URL")

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Erreur lors de la connexion à la base de données:", err)
	}
	defer db.Close()

	// Verify the connection by attempting a simple query
	if err := db.Ping(); err != nil {
		log.Fatalln("Erreur lors de la vérification de la connexion à la base de données:", err)
	}

	queries := New(db)
	return queries
}
