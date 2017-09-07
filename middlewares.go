package ethereal

import (
	"encoding/json"
	"github.com/justinas/alice"
	"log"
	"net/http"
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
	jwt            EtherealClaims
	statusError    int
	responseError  string
	authenticated  bool
	responseWriter http.ResponseWriter
	included       bool // flag is enabled or disabled authJwtToken
}

func (m middlewareJWTToken) Add(where *[]alice.Constructor, application *Application) {
	confToken := config("AUTH.JWT_TOKEN").(string)

	if confToken == "local" {
		m.included = true
		*where = append(*where, func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				m.responseWriter = w
				if check, err := m.jwt.Verify(r); !check {
					m.responseError = handlerErrorToken(err).Error()
				} else{
					m.authenticated = true
				}

				ctxStruct(application, m)
				handler.ServeHTTP(w, r)
			})
		})
	} else if confToken == "global" {
		// check jwt token all queries..
		*where = append(*where, func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if check, err := m.jwt.Verify(r); !check {
					m.responseError = handlerErrorToken(err).Error()
				}
				w.WriteHeader(m.statusError)
				json.NewEncoder(w).Encode(http.StatusText(m.statusError))
			})
		})
	} else {
		log.Println("Our config parameter AUTH.JWT_TOKEN = " + confToken)
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
