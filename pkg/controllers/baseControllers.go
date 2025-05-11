package controllers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/views/base"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := context.Background()
	if r.Method == http.MethodPost {
		err = r.ParseForm()
		if err != nil {
			log.SetOutput(os.Stdout)
			log.Println("Erreur lors du traitement du formulaire dans le controlleur connexion", err)
			return
		}
		email := r.Form.Get("email")
		password := r.Form.Get("mot de passe")
		db := database.Connection()
		defer db.Close()
		user, err := db.GetUserByEmail(ctx, email)
		if err != nil {
			log.SetOutput(os.Stdout)
			log.Println(err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
		if err != nil {
			log.SetOutput(os.Stdout)
			log.Println("erreur lors de la création du hash du mot de passe", err)
			return
		}
	}
	template := base.Login()
	err = template.Render(ctx, w)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println(err)
		return
	}

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := context.Background()
	if r.Method == http.MethodPost {
		err = r.ParseForm()
		if err != nil {
			log.SetOutput(os.Stdout)
			log.Println(err)
			return
		}
		nom := r.Form.Get("nom")
		postnom := r.Form.Get("postnom")
		prenom := r.Form.Get("prenom")
		email := r.Form.Get("email")
		password := r.Form.Get("mot de passe")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			log.Println("Erreur lors de la création du hash", err)
			return
		}
		db := database.Connection()
		defer db.Close()
		db.NewUser(ctx, database.NewUserParams{
			Nom:          sql.NullString{String: nom, Valid: true},
			Postnom:      sql.NullString{String: postnom, Valid: true},
			Prenom:       sql.NullString{String: prenom, Valid: true},
			Email:        email,
			PasswordHash: sql.NullString{String: string(hash), Valid: true},
		})
	}
	template := base.SignIn()
	err = template.Render(ctx, w)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println("Erreur lros de la création du template login", err)
		return
	}
}
