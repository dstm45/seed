package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/dstm45/seed/pkg/database"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var secret []byte = []byte(os.Getenv("SECRET"))

func BuildToken(user database.User) *http.Cookie {
	claims := TokenClaim{
		Email: user.Email.String,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "seed",
			Subject:   "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		AfficherErreur("Erreur lors de la création du token jwt", err)
	}
	cookie := http.Cookie{
		Name:    "auth",
		Value:   ss,
		Expires: claims.ExpiresAt.Add(time.Hour * 0),
	}
	return &cookie
}

func ParseToken(r *http.Request, email string) bool {
	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		AfficherErreur("le cookie est inéxistant", err)
		return false
	}
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		AfficherErreur("Erreur lors du parsing du token", err)
		return false
	}

	claims, ok := token.Claims.(*TokenClaim)
	if !ok || !token.Valid {
		return false
	}
	if claims.Email != email {
		return false
	}
	return true
}
