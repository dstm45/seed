package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dstm45/seed/pkg/database"
	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte(os.Getenv("SECRET"))

func BuildToken(user database.User) *http.Cookie {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "seed",
		Subject:   "auth",
		ID:        fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		log.Println("Erreur lors de la création du token jwt", err)
	}
	cookie := http.Cookie{
		Name:    "auth",
		Value:   ss,
		Expires: claims.ExpiresAt.Add(time.Hour * 0),
	}
	return &cookie
}

func ParseToken(r *http.Request) bool {
	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		log.Println("le cookie est inéxistant")
		return false
	}
	var tokenString = cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println("Erreur lors du parsing du token", err)
		return false
	}

	switch {
	case token.Valid:
		return true
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
		return false
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		fmt.Println("Invalid signature")
		return false
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		fmt.Println("Timing is everything")
		return false
	default:
		fmt.Println("Couldn't handle this token:", err)
	}
	return true
}
