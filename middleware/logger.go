package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("ami middleware, ami age print hobo")

		next.ServeHTTP(w, r)

		log.Println("ami middleware, ami pore print hobo")

		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
