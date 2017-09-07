package ethereal

import (
	"encoding/json"
	"github.com/justinas/alice"
	"net/http"
	"os"
	"strings"
)

/**
/ Add middleware in App under certain condition..
*/
type AddMiddleware interface {
	Add(*[]alice.Constructor, *Application)
}

type Middleware struct {
	// all middleware
	allMiddleware []AddMiddleware
	// middleware only included in application
	includeMiddleware []alice.Constructor
}

func (m Middleware) AddMiddleware(middleware ...AddMiddleware) {
	m.allMiddleware = append(m.allMiddleware, middleware...)
}

// Method loading middleware for application
func (m *Middleware) LoadApplication(application *Application) []alice.Constructor {
	for _, middleware := range m.allMiddleware {
		middleware.Add(&m.includeMiddleware, application)
	}
	return m.includeMiddleware
}

/**
/ ability to set jwt token all queries or choose query
*/
type middlewareJWTToken struct {
	status        int
	responseText      string
	authenticated bool
	responseWriter http.ResponseWriter
}

func (m middlewareJWTToken) Add(where *[]alice.Constructor, application *Application) {
	if os.Getenv("AUTH_JWT_TOKEN") != "" && os.Getenv("AUTH_JWT_TOKEN") == "global" {
		*where = append(*where, func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				headerBearer := r.Header.Get("Authorization")

				// get token
				if strings.HasPrefix(headerBearer, "Bearer") {
					token := strings.Replace(headerBearer, "Bearer", "", 1)
					token = strings.Trim(token, " ")

					if t, err := compareToken(token); err != nil && !t.Valid {
						w.WriteHeader(http.StatusNetworkAuthenticationRequired)
						json.NewEncoder(w).Encode(handlerErrorToken(err).Error())
						return
					}
				} else {
					// required authentication..
					//w.WriteHeader(http.StatusNetworkAuthenticationRequired)
					//json.NewEncoder(w).Encode(http.StatusText(http.StatusNetworkAuthenticationRequired))

					m.responseText = http.StatusText(http.StatusNetworkAuthenticationRequired)
					m.status = http.StatusNetworkAuthenticationRequired
					m.authenticated = false
					m.responseWriter = w
					ctxStruct(application, m)
					handler.ServeHTTP(w, r)
				}
			})
		})
	}
}

// ---- waiting for your implementation ------

// middleware set Accept-Language
func middlewareLocal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO Pipline choose
		// TODO set locale from request
		//app.Locale = parserLocale(r.Header["Accept-Language"])
		next.ServeHTTP(w, r)
	})
}
