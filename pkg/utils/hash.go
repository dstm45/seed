package utils

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/dstm45/seed/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Hasher(motDePasse string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(motDePasse), bcrypt.DefaultCost)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println("Erreur lors de la création du hash", err)
		return ""
	}
	return string(hash)
}

func ComparerHash(motDePasse, email string) bool {
	ctx := context.Background()
	db := database.Connection()
	hash, err := db.GetPasswordHash(ctx, sql.NullString{String: email, Valid: true})
	if err != nil {
		AfficherErreur("Erreur lors du retrait du hash dans la base de donnée", err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash.String), []byte(motDePasse))
	return err != nil
}
