package jwt

import "github.com/dgrijalva/jwt-go"

type ClaimsType struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
