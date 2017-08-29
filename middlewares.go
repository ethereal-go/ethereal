package ethereal

import (
	"net/http"
)

func middlewareLocal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Pipline choose Language
		app.Locale = parserLocale(r.Header["Accept-Language"])
		graphQL().fill()
		next.ServeHTTP(w, r)
	})
}
