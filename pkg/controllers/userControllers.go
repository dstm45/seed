package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dstm45/seed/pkg/views/user"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	template := user.Index(username)
	ctx := context.Background()
	template.Render(ctx, w)
}

func UserAnnonces(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	template := user.Annonces(username)
	ctx := context.Background()
	template.Render(ctx, w)
}

func UserDashboard(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	template := user.Dashboard(username)
	ctx := context.Background()
	template.Render(ctx, w)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println(err)
		return
	}
}
