package main

import (
	"github.com/dstm45/seed/pkg/config"
	"github.com/dstm45/seed/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.AfficherErreur("le fichier .env n'est pas présent", err)
		return
	}
	server := config.NewServer()
	server.Start()
}
