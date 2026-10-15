package middleware

import (
	"crypto/subtle"
	"log"
	"net/http"
	"strings"
)

func Authenticate(expectedToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scheme, token, ok := strings.Cut(r.Header.Get("Authorization"), " ")
			if !ok || !strings.EqualFold(scheme, "Bearer") ||
				subtle.ConstantTimeCompare([]byte(token), []byte(expectedToken)) != 1 {
				log.Println(r.URL.Path, "unauthorized")
				w.Header().Set("WWW-Authenticate", "Bearer")
				http.Error(w, "unauthorized request", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
