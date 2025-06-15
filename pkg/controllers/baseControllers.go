package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/authentication"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template := authentication.SignIn()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page signin", err)
			return
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement des formulaires dans la fonction signin", err)
			return
		}
		email := r.Form.Get("email")
		password := r.Form.Get("mot_de_passe")
		tokenValide := utils.ParseToken(r)
		passwordCorrect := utils.ComparerHash(password, email)
		if validiteToken && passwordCorrect{
			
		} else if !{			
			fmt.Println("email ou mot de passe incorrect")
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template := authentication.SignUp()
		err := template.Render(r.Context(), w)
		if err != nil {
			log.Println("Erreur lors du rendu de la page signup")
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur dans le traitement des formulaire dans la fonction signup", err)
			return
		}
		hashMdp := utils.Hasher(r.Form.Get("mot_de_passe"))
		user := database.NewUserParams{
			Nom:          sql.NullString{String: r.Form.Get("nom")},
			Prenom:       sql.NullString{String: r.Form.Get("prenom")},
			Email:        sql.NullString{String: r.Form.Get("email")},
			PasswordHash: sql.NullString{String: hashMdp},
		}
		db := database.Connection()
		err = db.NewUser(r.Context(), user)
		if err != nil {
			utils.AfficherErreur("Erreur lors de l'enregistrement de l'utilisateur dans la base de donnée", err)
			return
		}
		u, err := db.GetUserByEmail(r.Context(), sql.NullString{String: r.Form.Get("email"), Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors du retrait de l'utilisateur dans la base de donnée", err)
			return
		}
		http.SetCookie(w, utils.BuildToken(u))
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}
