package ethereal

import (
	"log"
	"net/http"
)

func middlewareLocal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header["Accept-Language"])
		// TODO Pipline choose Language
		parserLocale(r.Header["Accept-Language"])
		next.ServeHTTP(w, r)
	})
}
