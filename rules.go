package ethereal

import (
	"github.com/graphql-go/graphql"
)

type Rule interface {
	Verify(support interface{}) bool
}

type JwtTokenRule struct {
	exclude []*graphql.Object
	authenticated bool
}

func (jwt JwtTokenRule) Verify(support interface{}) bool {
	for _, oneType := range jwt.exclude {
		if oneType.Name() == support.(string) {
			return true
		}
	}
	return false
}
