// Package utils contient des fonctions utilitaires pour l'application.
package utils

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func Render(ctx context.Context, component templ.Component, w http.ResponseWriter) {
	err := component.Render(ctx, w)
	if err != nil {
		log.SetOutput(os.Stdout)
		log.Println(err)
		return
	}
}
