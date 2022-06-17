package jwt

import "github.com/golang-jwt/jwt"

type ClaimsType struct {
	Username       string `json:"username"`
	StandardClaims *jwt.StandardClaims
}
