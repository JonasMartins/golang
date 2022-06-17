package auth

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

// source: https://gqlgen.com/recipes/authentication/
// https://www.sohamkamani.com/golang/jwt-authentication/
// https://stackoverflow.com/questions/66090686/gqlgen-set-cookie-from-resolver
// https://github.com/99designs/gqlgen/issues/567

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	Name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

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
			if err != nil {
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
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

func validateAndGetUserID(c *http.Cookie) (string, error) {
	return "", nil
}

func getUserByID(db *gorm.DB, userId string) (*User, error) {
	user := &User{
		Name: "User",
	}
	return user, nil
}
