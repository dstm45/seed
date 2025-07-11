// Package utils contient des fonctions utilitaires pour l'application.
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
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

var secret []byte = []byte(os.Getenv("SECRET"))

func BuildToken(user database.User) *http.Cookie {
	claims := TokenClaim{
		Email: user.Email.String,
		Role:  user.TypeCompte.String,
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

func ParseToken(r *http.Request) (*TokenClaim, bool) {
	cookie, err := r.Cookie("auth")
	if err != nil {
		AfficherErreur("le cookie est inéxistant", err)
		return nil, false
	}
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		AfficherErreur("Erreur lors du parsing du token", err)
		return nil, false
	}
	claims, ok := token.Claims.(*TokenClaim)
	if !ok {
		AfficherErreur("Impossible de convertir les claims du token", nil)
		return nil, false
	}
	if !token.Valid {
		AfficherErreur("Token JWT invalide", nil)
		return nil, false
	}
	return claims, true
}

func DecodeToken(r *http.Request) *TokenClaim {
	cookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		AfficherErreur("Cookie Inexistant chez le client fonction decodetoken", err)
		return nil
	}
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		AfficherErreur("Erreur lors du parsing du token", err)
		return nil
	}
	claims, ok := token.Claims.(*TokenClaim)
	if !ok {
		AfficherErreur("Impossible de convertir les claims du token", nil)
		return nil
	}
	if !token.Valid {
		AfficherErreur("Token JWT invalide", nil)
		return nil
	}
	return claims
}

func DeleteToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
}
