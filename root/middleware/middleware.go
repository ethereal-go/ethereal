package middleware

import(
	"github.com/justinas/alice"
	"github.com/ethereal-go/ethereal/root/app"
	"net/http"
)

/**
/ Add middleware in App under certain condition..
*/
type AddMiddleware interface {
	Add(*[]alice.Constructor, *app.Application)
}

type Middleware struct {
	// all middleware
	AllMiddleware []AddMiddleware
	// middleware only included in application
	IncludeMiddleware []alice.Constructor

}

// Method loading middleware for application
func (m Middleware) LoadApplication(application *app.Application) []alice.Constructor {
	for _, middleware := range m.AllMiddleware {
		middleware.Add(&m.IncludeMiddleware, application)
	}
	return m.IncludeMiddleware
}

func (m Middleware) GetHandler(h http.HandlerFunc) http.Handler {
	return alice.New(m.IncludeMiddleware...).Then(h)
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