package ethereal

import (
	"github.com/graphql-go/graphql"
	"os"
)

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.RoleType.id")),
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.RoleType.name")),
		},
		"display_name": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.RoleType.display_name")),
		},
		"description": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.RoleType.description")),
		},
	},
})


var RoleField = graphql.Field{
	Type:        graphql.NewList(roleType),
	Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.Role.Description")),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var roles []Role
		app.Db.Find(&roles)

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
