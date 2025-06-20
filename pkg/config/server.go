package config

import (
	"net/http"

	"github.com/dstm45/seed/pkg/controllers"
	"github.com/dstm45/seed/pkg/middlewares"
	"github.com/dstm45/seed/pkg/utils"
)

type Myserver struct {
	*http.Server
}

var (
	index = middlewares.IsAuthenticated(controllers.UserIndex)
)

func NewServer() *Myserver {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /index", index)
	mux.HandleFunc("GET /signup", controllers.SignUp)
	mux.HandleFunc("POST /signup", controllers.SignUp)
	mux.HandleFunc("GET /signin", controllers.SignIn)
	mux.HandleFunc("POST /signin", controllers.SignIn)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("pkg/views/static"))))
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
