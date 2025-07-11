// Package utils contient des fonctions utilitaires pour l'application.
package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func SaveFile(w http.ResponseWriter, r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		AfficherErreur("Erreur lors du traitement du formulaire", err)
		return "", err
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		if err == http.ErrMissingFile {
			return "", nil
		}
		AfficherErreur("Erreur d'ouverture du fichier:", err)
		return "", err
	}
	defer file.Close()

	// Generate a unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)

	// Server-side and web-accessible paths
	saveDir := "pkg/views/static/images/profil"
	savePath := filepath.Join(saveDir, filename)
	webPath := filepath.Join("/static/images/profil", filename)

	// Save the file
	out, err := os.Create(savePath)
	if err != nil {
		log.Println("Erreur de création de fichier:", err)
		http.Error(w, "Erreur de sauvegarde", http.StatusInternalServerError)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Erreur lors de la copie du fichier:", err)
		http.Error(w, "Erreur de transfert de fichier", http.StatusInternalServerError)
		return "", err
	}
	return webPath, nil
}