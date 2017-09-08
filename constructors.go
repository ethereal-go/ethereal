package ethereal

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
)

// Here all constructors application, which return some structure...

func ConstructorI18N() *i18n.I18n {
	if App.I18n == nil {
		App.I18n = i18n.New(
			database.New(ConstructorDb()),
		)
	}
	return App.I18n
}

func ConstructorDb() *gorm.DB {
	if App.Db == nil {
		ConstructorConfig()
		App.Db = Database()
	}
	return App.Db

}

func ConstructorMiddleware() *Middleware {
	if App.Middleware == nil {
		App.Middleware = &Middleware{
			allMiddleware: []AddMiddleware{
				//list standard middleware
				//middlewareJWTToken{
				//	responseError: http.StatusText(http.StatusNetworkAuthenticationRequired),
				//	statusError:   http.StatusNetworkAuthenticationRequired,
				//},
			},
		}
	}
	return App.Middleware
}

func ConstructorConfig() *Config {
	if App.Config == nil {
		App.Config = &Config{}
		App.Config.LoadConfigFromApp()
	}
	return App.Config
}

//  init mutation global
func Mutations() GraphQlMutations {
	if mutations == nil {
		mutations = make(GraphQlMutations)
	}
	return mutations
}

//  init query global
func Queries() GraphQlQueries {
	if queries == nil {
		queries = make(GraphQlQueries)
	}
	return queries
}

// Function add default field mutation
func startMutations() map[string]*graphql.Field {
	Mutations().Add("createUser", &createUser)
	return mutations
}

func startQueries() map[string]*graphql.Field {
	Queries().Add("users", &UserField).Add("role", &RoleField)
	return queries
}
