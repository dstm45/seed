package utils

import (
	"log"
	"os"
)

func AfficherErreur(message string, err error) {
	log.SetOutput(os.Stdin)
	log.Println(message, err)
}
