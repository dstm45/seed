package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *Queries {
	// Open the database connection
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Erreur lors de la connexion à la base de données", err)
	}

	// Verify the connection by attempting a simple query
	if err := db.Ping(); err != nil {
		log.Fatalln("Erreur lors de la vérification de la connexion à la base de données", err)
	}

	connection := New(db)
	return connection
}
