package config

import (
	"net/http"

	"github.com/dstm45/seed/pkg/controllers"
	"github.com/dstm45/seed/pkg/database"
	"github.com/dstm45/seed/pkg/middlewares"
)

func NewRoutes(db *database.Queries) *http.ServeMux {
	mux := http.NewServeMux()
	user := controllers.UserController{
		DB: db,
	}
	event := controllers.EventController{
		DB: db,
	}
	admin := controllers.AdminController{
		DB: db,
	}
	auth := controllers.AuthController{
		DB: db,
	}
	index := middlewares.IsAuthenticated(user.Index)
	editProfile := middlewares.IsAuthenticated(user.Editer)
	editPassword := middlewares.IsAuthenticated(user.ChangePassword)
	flux := middlewares.IsAuthenticated(event.Flux)
	detailEvenement := middlewares.IsAuthenticated(event.Detail)

	settings := middlewares.IsAdmin(admin.Settings)
	dashboard := middlewares.IsAdmin(admin.Dashboard)
	manageUsers := middlewares.IsAdmin(admin.Users)
	manageEvent := middlewares.IsAdmin(admin.Events)
	encours := middlewares.IsAdmin(admin.Encours)

	mux.HandleFunc("GET /signup", auth.SignUp)
	mux.HandleFunc("POST /signup", auth.SignUp)
	mux.HandleFunc("GET /signin", auth.SignIn)
	mux.HandleFunc("POST /signin", auth.SignIn)
	mux.HandleFunc("GET /logout", auth.Logout)

	mux.HandleFunc("/index", index)
	mux.HandleFunc("/", index)
	mux.HandleFunc("GET /editprofile", editProfile)
	mux.HandleFunc("GET /editpassword", editPassword)
	mux.HandleFunc("GET /evenement/flux", flux)
	mux.HandleFunc("GET /evenement/detail/{uuid}", detailEvenement)

	mux.HandleFunc("GET /admin/dashboard", dashboard)
	mux.HandleFunc("GET /admin/settings", settings)
	mux.HandleFunc("GET /admin/encours", encours)
	mux.HandleFunc("GET /admin/manageusers", manageUsers)
	mux.HandleFunc("GET /admin/manageevent", manageEvent)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("pkg/views/static"))))
	return mux
}
