package graphQL

import (
	"fmt"
	"github.com/agoalofalife/ethereal"
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
	Description: "Get single todo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		var users []ethereal.User
		var role ethereal.Role
		ethereal.A.Db.Find(&users)

		idQuery, isOK := params.Args["id"].(string)

		if isOK {
			for _, user := range users {
				if strconv.Itoa(int(user.ID)) == idQuery {
					ethereal.A.Db.Model(&user).Related(&role)
					user.Role = role
					fmt.Println(user.Role)
					return []ethereal.User{user}, nil
				}
			}
		}
		return users, nil
	},
}
