package controllers

import (
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/authentication"
	"github.com/dstm45/seed/pkg/views/component"
	"github.com/jackc/pgx/v5/pgtype"
)

// AuthController holds the dependencies for the authentication handlers.
type AuthController struct {
	DB *database.Queries
}

// SignIn handles user login.
func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := authentication.SignIn()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page signin", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement des formulaires dans la fonction signin", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		user, err := c.DB.GetUserByEmail(r.Context(), pgtype.Text{String: email, Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors du retrait de l'utilisateur dans la fonction signin", err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		if !utils.ComparerHash(password, user.PasswordHash.String) {
			// In a real application, you might want to show a more generic error message
			// to avoid revealing whether an email address is registered.
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		http.SetCookie(w, utils.BuildToken(user))
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// SignUp handles new user registration.
func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := authentication.SignUp()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page signup", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur dans le traitement des formulaire dans la fonction signup", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		hashMdp := utils.Hasher(r.Form.Get("mot_de_passe"))
		userParams := database.NewUserParams{
			Nom:          pgtype.Text{String: r.Form.Get("nom"), Valid: true},
			Prenom:       pgtype.Text{String: r.Form.Get("prenom"), Valid: true},
			Pseudonyme:   pgtype.Text{String: r.Form.Get("pseudonyme"), Valid: true},
			Email:        pgtype.Text{String: r.Form.Get("email"), Valid: true},
			PasswordHash: pgtype.Text{String: hashMdp, Valid: true},
		}

		user, err := c.DB.NewUser(r.Context(), userParams)
		if err != nil {
			// This could be a unique constraint violation (e.g., email already exists)
			utils.AfficherErreur("Erreur lors de l'enregistrement de l'utilisateur dans la base de donnée", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, utils.BuildToken(user))
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// Logout handles user logout.
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		utils.DeleteToken(w)
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
