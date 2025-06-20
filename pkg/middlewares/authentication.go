package middlewares

import (
	"net/http"

	"github.com/dstm45/seed/pkg/utils"
)

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.ParseToken(r) {
			next(w, r)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}
}
