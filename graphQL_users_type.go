package ethereal

import (
	"github.com/graphql-go/graphql"
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
			Description: "",
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"role": &graphql.Field{
			Type: roleType,
		},
	},
})

var UserField = graphql.Field{
	Type:        graphql.NewList(usersType),
	Description: string(ConstructorI18N()().T("en-US", "graphQL.User.Description")),
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
