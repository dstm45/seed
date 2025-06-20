package controllers

import (
	"fmt"
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/authentication"
	"github.com/jackc/pgx/v5/pgtype"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	conn := database.Connection()
	db := database.New(conn)
	defer conn.Close(r.Context())
	if r.Method == http.MethodGet {
		template := authentication.SignIn()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page signin", err)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement des formulaires dans la fonction signin", err)
			return
		}
		email := r.Form.Get("email")
		password := r.Form.Get("mot_de_passe")
		user, err := db.GetUserByEmail(r.Context(), pgtype.Text{String: email, Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors du retrait de l'utilisateur dans la fonction signin", err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		// Compare the password with the hash from the user record
		if !utils.ComparerHash(password, user.PasswordHash.String) {
			fmt.Println("email ou mot de passe incorrect")
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		// Build and set the token cookie
		http.SetCookie(w, utils.BuildToken(user))
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	conn := database.Connection()
	db := database.New(conn)
	defer conn.Close(r.Context())
	if r.Method == http.MethodGet {
		template := authentication.SignUp()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page signup", err)
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur dans le traitement des formulaire dans la fonction signup", err)
			return
		}
		hashMdp := utils.Hasher(r.Form.Get("mot_de_passe"))
		user := database.NewUserParams{
			Nom:          pgtype.Text{String: r.Form.Get("nom"), Valid: true},
			Prenom:       pgtype.Text{String: r.Form.Get("prenom"), Valid: true},
			Email:        pgtype.Text{String: r.Form.Get("email"), Valid: true},
			PasswordHash: pgtype.Text{String: hashMdp, Valid: true},
		}
		err = db.NewUser(r.Context(), user)
		if err != nil {
			utils.AfficherErreur("Erreur lors de l'enregistrement de l'utilisateur dans la base de donnée", err)
			return
		}
		u, err := db.GetUserByEmail(r.Context(), pgtype.Text{String: r.Form.Get("email"), Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors du retrait de l'utilisateur dans la base de donnée", err)
			return
		}
		http.SetCookie(w, utils.BuildToken(u))
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteToken(w)
	http.Redirect(w, r, "/signup", http.StatusSeeOther)
}
