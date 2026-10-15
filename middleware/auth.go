package middleware

import (
	"log"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, "authenticating")
		next.ServeHTTP(w, r)
	})
}
