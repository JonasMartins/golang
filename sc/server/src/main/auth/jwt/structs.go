package jwt

import "github.com/golang-jwt/jwt"

type ClaimsType struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
