package ethereal

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
	"os"
	"strings"
)

/**
/ Add middleware in App under certain condition..
*/
type AddMiddleware interface {
	Add(*[]alice.Constructor)
}

type Middleware struct {
	// all middleware
	allMiddleware []AddMiddleware
	// middleware only included in application
	includeMiddleware []alice.Constructor
}

func ConstructorMiddleware() *Middleware {
	if app.Middleware == nil {
		app.Middleware = &Middleware{allMiddleware: []AddMiddleware{
			// list standard
			middlewareJWTToken{},
		}}
	}

	return app.Middleware
}
func (m Middleware) AddMiddleware(middleware ...AddMiddleware) {
	m.allMiddleware = append(m.allMiddleware, middleware...)
}

// Method loading middleware for application
func (m *Middleware) LoadApplication() []alice.Constructor {
	for _, middleware := range m.allMiddleware {
		middleware.Add(&m.includeMiddleware)
	}
	return m.includeMiddleware
}

type middlewareJWTToken struct{}

func (m middlewareJWTToken) Add(where *[]alice.Constructor) {
	if os.Getenv("AUTH_JWT_TOKEN") != "" && os.Getenv("AUTH_JWT_TOKEN") == "true" {
		*where = append(*where, func(handler http.Handler) http.Handler {
			// To add the ability to select the type of authenticate
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")

				//app.GraphQlMutation["createJWTToken"] = &createJWTToken

				// get token
				if strings.HasPrefix(authHeader, "Bearer") {
					token := strings.Replace(authHeader, "Bearer", "", 1)
					token = strings.Trim(token, " ")

					if t, err := compareToken(token); err != nil && !t.Valid {
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
