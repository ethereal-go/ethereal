package ethereal

import (
	"github.com/graphql-go/graphql"
	"fmt"
)

type Rule interface {
	Verify() bool
}

type JwtTokenRule struct {
	exclude []*graphql.Object
}

func (jwt JwtTokenRule) Verify() bool {
	fmt.Println("verify..")
	return true
}