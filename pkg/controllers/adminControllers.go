package controllers

import (
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/admin"
	"github.com/dstm45/seed/pkg/views/component"
)

// AdminController holds the dependencies for the admin handlers.
type AdminController struct {
	DB *database.Queries
}

// Dashboard handles the rendering of the admin dashboard page.
func (c *AdminController) Dashboard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := admin.Dashboard()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la fonction Dashboard", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// ManageUsers handles the rendering of the user management page.
func (c *AdminController) ManageUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := admin.ManageUsers()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la fonction ManageUsers", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// ManageEvents handles the rendering of the event management page.
func (c *AdminController) ManageEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := admin.ManageEvent()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la fonction ManageEvent", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// Settings handles the rendering of the admin settings page.
func (c *AdminController) Settings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := admin.Settings()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la fonction settings", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}

// Encours handles the rendering of the "in progress" page.
func (c *AdminController) Encours(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := admin.EnCours()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la fonction en cours", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		component.MethodNotAllowed().Render(r.Context(), w)
	}
}
