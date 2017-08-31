package ethereal

import (
	"fmt"
	"net/http"
	"strings"
)

// middleware set Accept-Language
func middlewareLocal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Pipline choose
		// TODO set locale from request
		//app.Locale = parserLocale(r.Header["Accept-Language"])
		next.ServeHTTP(w, r)
	})
}

// To add the ability to select the type of authenticate
func middlewareAuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// get token
		if strings.HasPrefix(authHeader, "Bearer") {
			token := strings.Replace(authHeader, "Bearer", "", 1)
			token = strings.Trim(token, " ")

			if t, err := compareToken(token); err == nil && t.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusNetworkAuthenticationRequired)
				fmt.Fprint(w, handlerErrorToken(err).Error())
				return
			}

		} else {
			// required authentication..
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			fmt.Fprint(w, http.StatusText(http.StatusNetworkAuthenticationRequired))
			return
		}
	})
}
