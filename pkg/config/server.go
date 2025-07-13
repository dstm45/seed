// Package config fournit la configuration et l'initialisation du serveur.
package config

import (
	"net/http"

	"github.com/dstm45/seed/pkg/controllers"
	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/middlewares"
	"github.com/dstm45/seed/pkg/utils"
)

type Myserver struct {
	*http.Server
}

func NewServer() *Myserver {
	conn := database.Connection()
	db := database.New(conn)
	// controllers
	user := controllers.UserController{
		DB: db,
	}
	admin := controllers.AdminController{
		DB: db,
	}
	auth := controllers.AuthController{
		DB: db,
	}
	index := middlewares.IsAuthenticated(user.Index)
	editprofile := middlewares.IsAuthenticated(user.EditProfile)
	editpassword := middlewares.IsAuthenticated(user.EditPassword)
	evenement := middlewares.IsAuthenticated(user.Event)

	settings := middlewares.IsAdmin(admin.Settings)
	dashboard := middlewares.IsAdmin(admin.Dashboard)
	manageusers := middlewares.IsAdmin(admin.ManageUsers)
	manageevent := middlewares.IsAdmin(admin.ManageEvents)
	encours := middlewares.IsAdmin(admin.Encours)

	mux := http.NewServeMux()
	// auth
	mux.HandleFunc("GET /signup", auth.SignUp)
	mux.HandleFunc("POST /signup", auth.SignUp)
	mux.HandleFunc("GET /signin", auth.SignIn)
	mux.HandleFunc("POST /signin", auth.SignIn)
	mux.HandleFunc("GET /logout", auth.Logout)
	// user
	mux.HandleFunc("GET /index", index)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/editprofile", editprofile)
	mux.HandleFunc("/editpassword", editpassword)
	mux.HandleFunc("/liste", user.ShowEvent)
	mux.HandleFunc("GET /evenement/detail/{id}", evenement)
	// admin
	mux.HandleFunc("GET /admin/dashboard", dashboard)
	mux.HandleFunc("GET /admin/settings", settings)
	mux.HandleFunc("GET /admin/encours", encours)
	mux.HandleFunc("GET /admin/manageusers", manageusers)
	mux.HandleFunc("GET /admin/manageevent", manageevent)
	// fileserver
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
