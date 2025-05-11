package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dstm45/seed/pkg/database"
	"github.com/golang-jwt/jwt/v5"
)

func BuildToken(user *database.User) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	fmt.Println(tokenString, err)
}
