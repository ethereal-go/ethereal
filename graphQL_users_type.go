package ethereal

import (
	"github.com/graphql-go/graphql"
	"os"
	"strconv"
)

/**
/ User Type
*/
var usersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.UserType.id")),
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.UserType.email")),
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.UserType.name")),
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.UserType.password")),
		},
		"role": &graphql.Field{
			Type:        roleType,
			Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.UserType.role")),
		},
	},
})

var UserField = graphql.Field{
	Type:        graphql.NewList(usersType),
	Description: string(ConstructorI18N().T(os.Getenv("LOCALE"), "graphQL.User.Description")),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		var users []*User
		app.Db.Find(&users)

		idQuery, isOK := params.Args["id"].(string)

		if isOK {
			for _, user := range users {
				if strconv.Itoa(int(user.ID)) == idQuery {
					var role Role
					app.Db.Model(&user).Related(&role)
					user.Role = role
					return []User{*user}, nil
				}
			}
		}

		for _, user := range users {
			var role Role
			app.Db.Model(&user).Related(&role)
			user.Role = role
		}

		return users, nil
	},
}
