package controllers

import (
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/dstm45/seed/pkg/views/component"
	"github.com/dstm45/seed/pkg/views/evenement"
	"github.com/google/uuid"
)

type EventController struct {
	DB *database.Queries
}

func (c EventController) Flux(w http.ResponseWriter, r *http.Request) {
	evenements, err := c.DB.GetAllEvent(r.Context())
	if err != nil {
		utils.AfficherErreur("Erreur lors du retrait des utilisateur dans la base de données dans la fonction show event", err)
		return
	}
	template := evenement.FluxEvenement(c.DB, evenements)
	err = template.Render(r.Context(), w)
	if err != nil {
		utils.AfficherErreur("Erreur lors du rendu de la fonction showevent", err)
		return
	}
}

func (c EventController) Detail(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		utils.AfficherErreur("Erreur lors de la recupération de l'uuid ", err)
		return
	}
	event, err := c.DB.GetEventByUUID(r.Context(), uuid)
	if err != nil {
		utils.AfficherErreur("Erreur lors du retrait de l'évenement dans la base de données", err)
		return
	}
	template := evenement.Evenement(event)
	template.Render(r.Context(), w)
}

func (c EventController) Editer(w http.ResponseWriter, r *http.Request) {}

func (c EventController) Supprimer(w http.ResponseWriter, r *http.Request) {}

func (c EventController) Ajouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template := evenement.AddEvent()
		err := template.Render(r.Context(), w)
		if err != nil {
			utils.AfficherErreur("Erreur lors du rendu de la page addevent", err)
			return
		}
	case http.MethodPost:
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.AfficherErreur("Erreur lors du traitement du formulaire dans la fonction addevent", err)
			return
		}
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
