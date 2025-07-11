// Package middlewares contient les middlewares pour l'authentification et l'autorisation.
package middlewares

import (
	"context"
	"net/http"

	"github.com/dstm45/seed/pkg/utils"
)

type contextKey string

const UserContextKey contextKey = "user"

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, valid := utils.ParseToken(r)
		if !valid {
			// effacer le cookie invalide
			http.SetCookie(w, &http.Cookie{
				Name:   "auth",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			})
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenClaim := utils.DecodeToken(r)
		role := tokenClaim.Role
		if role != "admin" {
			http.Redirect(w, r, "/503", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
