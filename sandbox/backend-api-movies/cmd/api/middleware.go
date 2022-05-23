package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type,Authorization")
		next.ServeHTTP(w, r)
	})
}

func (app *application) verifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

		}

		headerArr := strings.Split(authHeader, " ")
		if len(headerArr) != 2 {
			app.errorJSON(w, errors.New("invalid auth header "))
			return
		}
		if headerArr[0] != "Bearer" {
			app.errorJSON(w, errors.New("missing bearer"))
			return
		}

		token := headerArr[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))
		if err != nil {
			app.errorJSON(w, errors.New("unauthorized - failed hmac check"), http.StatusForbidden)
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("unauthorized - token expired"), http.StatusForbidden)
			return
		}

		if !claims.AcceptAudience("domain.com") {
			app.errorJSON(w, errors.New("unauthorized - invalid domain"), http.StatusForbidden)
			return
		}

		userID := claims.Subject
		if userID == "" {
			app.errorJSON(w, errors.New("unauthorized"), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
