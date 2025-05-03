package main

import (
	"log"
	"net/http"

	"github.com/dstm45/seed/pkg/controllers"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8888",
	}
	//user routes
	mux.HandleFunc("GET /{username}", controllers.UserIndex)
	mux.HandleFunc("GET /{username}/annonces", controllers.UserAnnonces)
	mux.HandleFunc("GET /{username}/dashboard", controllers.UserDashboard)
	log.Fatalln(server.ListenAndServe())
}
