package main

import (
	"log"

	"github.com/dstm45/seed/pkg/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	server := config.NewServer()
	server.Start()
}
