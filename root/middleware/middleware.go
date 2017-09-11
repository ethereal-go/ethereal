package middleware

import(
	"github.com/justinas/alice"
	"github.com/ethereal-go/ethereal/root"
)

/**
/ Add middleware in App under certain condition..
*/
type AddMiddleware interface {
	Add(*[]alice.Constructor, *root.Application)
}

type Middleware struct {
	// all middleware
	allMiddleware []AddMiddleware
	// middleware only included in application
	includeMiddleware []alice.Constructor
}
