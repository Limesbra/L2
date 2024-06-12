package middleware

import (
	"log"
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s, %s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}