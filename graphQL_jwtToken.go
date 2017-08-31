package ethereal

import (
	"github.com/agoalofalife/ethereal/utils"
	"github.com/graphql-go/graphql"
	"github.com/dgrijalva/jwt-go"

	"errors"
)

const (
	errorInputData = "Login or Password not valid"
)
var jwtType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JWTToken",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type:        graphql.String,
			Description: "",
		},
	},
})

//var jwtField = graphql.Field{
//	Type:        graphql.NewList(roleType),
//	Description: "",
//	Args: graphql.FieldConfigArgument{
//		"login": &graphql.ArgumentConfig{
//			Type: graphql.String,
//		},
//		"password": &graphql.ArgumentConfig{
//			Type: graphql.String,
//		},
//	},
//	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//		var roles []Role
//		app.Db.Find(&roles)
//
//		idQuery, isOK := params.Args["id"].(string)
//		if isOK {
//			for _, role := range roles {
//				if string(role.ID) == idQuery {
//					return role, nil
//				}
//			}
//		}
//		return roles, nil
//	},
//}

/**
/ Create Token
*/
var createJWTToken = graphql.Field{
	Type:        jwtType,
	Description: "Create new jwt-token",
	Args: graphql.FieldConfigArgument{
		"login": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var user User
		var generateToken string
		login, _ := params.Args["login"].(string)
		password, _ := params.Args["password"].(string)

		app.Db.Where("email = ?", login).First(&user)

		if utils.CompareHashPassword([]byte(user.Password), []byte(password)) {
			claims := EtherealClaims{
				jwt.StandardClaims{
					ExpiresAt: 15000,
					Issuer:    user.Email,
				},
			}
			// TODO add choose crypt via configuration!
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			generateToken, _ = token.SignedString(JWTKEY())

		} else{
			return nil, errors.New(errorInputData)
		}

		return 	struct {
			Token string `json:"token"`
		}{generateToken}, nil
	},
}
