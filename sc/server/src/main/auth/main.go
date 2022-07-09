package auth

import (
	"context"
	"net/http"
	"os"

	jwtLocal "src/main/auth/jwt"

	"github.com/golang-jwt/jwt"

	"gorm.io/gorm"
)

// source: https://gqlgen.com/recipes/authentication/
// https://www.sohamkamani.com/golang/jwt-authentication/
// https://stackoverflow.com/questions/66090686/gqlgen-set-cookie-from-resolver
// https://github.com/99designs/gqlgen/issues/567

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"id"}
var responseWriterCtxKey = &contextKey{"responseWriter"}
var jwtKey = []byte(jwtLocal.GetJwtSecret())

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			/*
				if r.Header["Origin"] != nil && r.Header["Origin"][0] == os.Getenv("ORIGIN_LOGIN") {
					ctx := context.WithValue(r.Context(), responseWriterCtxKey, w)
					r = r.WithContext(ctx)
					next.ServeHTTP(w, r)
					return
				} else { */
			c, err := r.Cookie(os.Getenv("COOKIE_NAME"))

			// Allow unauthenticated users in
			if err != nil || c == nil {
				//http.Error(w, "Not authorized", http.StatusForbidden)
				//return
				ctx := context.WithValue(r.Context(), responseWriterCtxKey, w)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			userId, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Not authorized", http.StatusForbidden)
				return
			}
			// put it in context
			rootCtx := context.WithValue(r.Context(), userCtxKey, userId)
			ctx := context.WithValue(rootCtx, responseWriterCtxKey, w)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			//}
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForUserIdContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}

// Finding the response writer inside context
func ForResponseWriterContext(ctx context.Context) http.ResponseWriter {
	raw, _ := ctx.Value(responseWriterCtxKey).(http.ResponseWriter)
	return raw
}

// Decodes cookie to find the token and validates its content
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

/*
func getUserByID(db *gorm.DB, userId string) (*user.User, error) {
	user := user.User{}
	if findUser := db.First(&user, "id = ?", userId); findUser.Error != nil {
		return nil, findUser.Error
	} else {
		return &user, nil
	}
}
*/
