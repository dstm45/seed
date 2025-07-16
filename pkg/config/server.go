// Package config fournit la configuration et l'initialisation du serveur.
package config

import (
	"net/http"

	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/utils"
)

type Myserver struct {
	*http.Server
}

func NewServer() *Myserver {
	conn := database.Connection()
	db := database.New(conn)
	// controllers
	mux := NewRoutes(db)
	newServer := Myserver{
		&http.Server{
			Handler: mux,
			Addr:    "0.0.0.0:8888",
		},
	}
	return &newServer
}

func (s Myserver) Start() {
	err := s.ListenAndServe()
	utils.AfficherErreur("Erreur lors du lancement du serveur", err)
}
