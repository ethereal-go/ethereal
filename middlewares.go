package ethereal

import (
	"github.com/qor/i18n"
	"log"
	"net/http"
)

func middlewareLocal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Pipline choose Language
		app.Locale = parserLocale(r.Header["Accept-Language"])
		next.ServeHTTP(w, r)
	})
}
