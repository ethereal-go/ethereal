package ethereal

import "github.com/graphql-go/graphql"

var mutations GraphQlMutations
var queries GraphQlQueries

type GraphQlMutations map[string]*graphql.Field
type GraphQlQueries map[string]*graphql.Field

type GraphFields interface {
	Add(name string, field *graphql.Field) GraphQlMutations
}

/**
/ Methods add new field (query or mutations) in GraphQl{types}..
*/
func (g GraphQlMutations) Add(name string, field *graphql.Field) GraphQlMutations {
	mutations[name] = field
	return g
}

func (g GraphQlQueries) Add(name string, field *graphql.Field) GraphQlQueries {
	queries[name] = field
	return g
}
