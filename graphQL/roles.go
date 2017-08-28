package graphQL

import (
	"github.com/agoalofalife/ethereal"
	"github.com/graphql-go/graphql"
)

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"display_name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RoleField = graphql.Field{
	Type:        graphql.NewList(roleType),
	Description: "Get single todo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var roles []ethereal.Role
		ethereal.A.Db.Find(&roles)

		idQuery, isOK := params.Args["id"].(string)
		if isOK {
			for _, role := range roles {
				if string(role.ID) == idQuery {
					return role, nil
				}
			}
		}

		return roles, nil
	},
}
