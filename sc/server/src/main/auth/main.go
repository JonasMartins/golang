package auth

import (
	"context"
	"net/http"

	jwtLocal "src/main/auth/jwt"

	"github.com/golang-jwt/jwt"

	"src/infra/orm/gorm/models/user"

	"gorm.io/gorm"
)

// source: https://gqlgen.com/recipes/authentication/
// https://www.sohamkamani.com/golang/jwt-authentication/
// https://stackoverflow.com/questions/66090686/gqlgen-set-cookie-from-resolver
// https://github.com/99designs/gqlgen/issues/567

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}
var jwtKey = []byte(jwtLocal.GetJwtSecret())

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("sc_cook")

			// Allow unauthenticated users in
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			userId, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// get the user from the database
			user, err := getUserByID(db, userId)
			if err != nil || user == nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *user.User {
	raw, _ := ctx.Value(userCtxKey).(*user.User)
	return raw
}

func validateAndGetUserID(c *http.Cookie) (string, error) {

	token := c.Value

	claims := &jwtLocal.ClaimsType{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "Unauthorized", err
		}
	}
	if !tkn.Valid {
		return "Unauthorized", err
	}
	return claims.Id, nil
}

func getUserByID(db *gorm.DB, userId string) (*user.User, error) {
	user := user.User{}
	if findUser := db.First(&user, "id = ?", userId); findUser.Error != nil {
		return nil, findUser.Error
	} else {
		return &user, nil
	}
}
