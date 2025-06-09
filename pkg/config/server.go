package config

import (
	"log"
	"net/http"

	"github.com/dstm45/seed/pkg/controllers"
)

type Myserver struct {
	*http.Server
}

func NewServer() *Myserver {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{username}", controllers.UserIndex)
	// mux.HandleFunc("GET /user/{username}/annonces", controllers.UserAnnonces)
	// mux.HandleFunc("GET /user/{username}/dashboard", controllers.UserDashboard)
	// mux.HandleFunc("GET /login", controllers.Login)
	// mux.HandleFunc("POST /login", controllers.Login)
	mux.HandleFunc("GET /connexion", controllers.SignIn)
	mux.HandleFunc("POST /connexion", controllers.SignIn)
	fileSystem := http.FileServer(http.Dir("../views/static"))
	mux.Handle("/", fileSystem)
	newServer := Myserver{
		&http.Server{
			Handler: mux,
			Addr:    "0.0.0.0:8888",
		},
	}
	return &newServer
}

func (s Myserver) Start() {
	log.Fatalln(s.ListenAndServe())
}
