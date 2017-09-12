package ethereal

import (
	"github.com/ethereal-go/ethereal/utils"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"github.com/ethereal-go/ethereal/root/app"
)

/**
/ User Type
*/
var usersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.UserType.id")),
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.UserType.email")),
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.UserType.name")),
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.UserType.password")),
		},
		"role": &graphql.Field{
			Type:        roleType,
			Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.UserType.role")),
		},
	},
})

/**
/ Create User
*/
var createUser = graphql.Field{
	Type:        usersType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		email, _ := params.Args["email"].(string)
		name, _ := params.Args["name"].(string)
		password, _ := params.Args["password"].(string)
		role, _ := params.Args["role"].(int)

		hashedPassword, err := utils.HashPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(`Error hash password create User service.`)
		}

		var user = User{Email: email, Name: name, Password: string(hashedPassword), RoleID: role}
		App.Db.Create(&user)

		return user, nil
	},
}

var UserField = graphql.Field{
	Type:        graphql.NewList(usersType),
	Description: string(ConstructorI18N().T(GetCnf("L18N.LOCALE").(string), "graphQL.User.Description")),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// code for local jwt token..

		//jwtAuth := params.Context.Value("middlewareJWTToken").(middlewareJWTToken)
		//
		//if jwtAuth.included == false || jwtAuth.authenticated {
		db := params.Context.Value("*Application").(*app.Application).Db
			var users []*User
		db.Find(&users)

			idQuery, isOK := params.Args["id"].(string)

			if isOK {
				for _, user := range users {
					if strconv.Itoa(int(user.ID)) == idQuery {
						var role Role
						db.Model(&user).Related(&role)
						user.Role = role
						return []User{*user}, nil
					}
				}
			}

			for _, user := range users {
				var role Role
				db.Model(&user).Related(&role)
				user.Role = role
			}

			return users, nil
		//}
		//jwtAuth.responseWriter.WriteHeader(jwtAuth.statusError)
		//json.NewEncoder(jwtAuth.responseWriter).Encode(http.StatusText(jwtAuth.statusError))
		//return nil, errors.New(jwtAuth.responseError)
	},
}
