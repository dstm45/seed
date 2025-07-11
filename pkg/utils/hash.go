// Package utils contient des fonctions utilitaires pour l'application.
package utils

import (
	"log"
	"os"

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

func ComparerHash(motDePasse, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(motDePasse))
	return err == nil
}
