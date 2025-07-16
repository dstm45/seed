// Package controllers contains all the controllers
package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/component"
	"github.com/dstm45/seed/pkg/views/user"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserController struct {
	DB *database.Queries
}

func (c UserController) Index(w http.ResponseWriter, r *http.Request) {
	claim := utils.DecodeToken(r)
	ctx := r.Context()
	switch r.Method {
	case http.MethodGet:
		utilisateur, err := c.DB.GetUserByEmail(ctx, pgtype.Text{String: claim.Email, Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors de l'accès de l'email dans la fonction index", err)
			return
		}
		evenements, err := c.DB.GetEventByUserEmail(r.Context(), pgtype.Text{String: claim.Email, Valid: true})
		if err != nil {
			utils.AfficherErreur("Erreur lors du retrait des évenements dans la base de données", err)
			return
		}
		template := user.Index(utilisateur, evenements)
		template.Render(ctx, w)
	}
}

func (c UserController) Editer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userInfo := utils.DecodeToken(r)
		utilisateur, err := c.DB.GetUserByEmail(r.Context(), pgtype.Text{String: userInfo.Email, Valid: true})
		if err != nil {
		}
		template := user.EditUser(utilisateur)
		template.Render(r.Context(), w)

	case http.MethodPost:
		var webPath string
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement du formulaire", err)
			http.Error(w, "Erreur lors du traitement du formulaire", http.StatusInternalServerError)
			return
		}

		file, handler, err := r.FormFile("avatar")
		if err != nil {
			if err != http.ErrMissingFile {
				utils.AfficherErreur("Erreur d'ouverture du fichier:", err)
				http.Error(w, "Erreur d'ouverture du fichier", http.StatusInternalServerError)
				return
			}
		} else {
			defer file.Close()

			filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
			saveDir := "pkg/views/static/images/profil"
			savePath := filepath.Join(saveDir, filename)
			webPath = filepath.Join("/static/images/profil", filename)

			out, err := os.Create(savePath)
			if err != nil {
				utils.AfficherErreur("Erreur de création de fichier:", err)
				http.Error(w, "Erreur de sauvegarde", http.StatusInternalServerError)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				utils.AfficherErreur("Erreur lors de la copie du fichier:", err)
				http.Error(w, "Erreur de transfert de fichier", http.StatusInternalServerError)
				return
			}
		}

		userInfo := utils.DecodeToken(r)

		if webPath == "" {
			currentUser, err := c.DB.GetUserByEmail(r.Context(), pgtype.Text{String: userInfo.Email, Valid: true})
			if err != nil {
				utils.AfficherErreur("Erreur lors de la récupération des informations utilisateur", err)
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
				return
			}
			webPath = currentUser.CheminPhoto.String
		}

		newInformations := database.UpdateDescriptionParams{
			Nom:         pgtype.Text{String: r.FormValue("nom"), Valid: true},
			Prenom:      pgtype.Text{String: r.FormValue("prenom"), Valid: true},
			Description: pgtype.Text{String: r.FormValue("description"), Valid: true},
			Pseudonyme:  pgtype.Text{String: r.FormValue("username"), Valid: true},
			CheminPhoto: pgtype.Text{String: webPath, Valid: true},
			Email:       pgtype.Text{String: userInfo.Email, Valid: true},
		}

		err = c.DB.UpdateDescription(r.Context(), newInformations)
		if err != nil {
			utils.AfficherErreur("Erreur mise à jour utilisateur", err)
			http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		template := component.MethodNotAllowed()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu dans la fontion controller", err)
			return
		}
	}
}

func (c UserController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := user.ChangePassword()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page editpassword", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement du formulaire dans la fonction editpassword", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		passwordHash := utils.Hasher(r.Form.Get("mot_de_passe"))
		parameters := database.UpdatePasswordParams{
			PasswordHash: pgtype.Text{String: passwordHash, Valid: true},
			Email:        pgtype.Text{String: utils.DecodeToken(r).Email, Valid: true},
		}
		err = c.DB.UpdatePassword(r.Context(), parameters)
		if err != nil {
			utils.AfficherErreur("Erreur lors de la mise à jour du mot de passe", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		template := component.MethodNotAllowed()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu dans la fontion controller", err)
			return
		}
	}
}
