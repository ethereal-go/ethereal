package ethereal
//
//import (
//	"github.com/justinas/alice"
//	"net/http"
//)

//
///**
/// Add middleware in App under certain condition..
//*/
//type AddMiddleware interface {
//	Add(*[]alice.Constructor, *Application)
//}
//
//type Middleware struct {
//	// all middleware
//	allMiddleware []AddMiddleware
//	// middleware only included in application
//	includeMiddleware []alice.Constructor
//}
//
//func (m *Middleware) AddMiddleware(middleware ...AddMiddleware) {
//	m.allMiddleware = append(m.allMiddleware, middleware...)
//}
//
//// Method loading middleware for application
//func (m Middleware) LoadApplication(application *Application) []alice.Constructor {
//	for _, middleware := range m.allMiddleware {
//		middleware.Add(&m.includeMiddleware, application)
//	}
//	return m.includeMiddleware
//}
//
//// ---- waiting for your implementation ------
//
//// middleware set Accept-Language
//func middlewareLocal(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// TODO Pipline choose
//		// TODO set locale from request
//		//app.Locale = parserLocale(r.Header["Accept-Language"])
//		next.ServeHTTP(w, r)
//	})
//}
