package ethereal

import (
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/ethereal-go/ethereal/root/config/json"
	"github.com/ethereal-go/ethereal/root/middleware"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
)

// Here all constructors application, which return some structure...

// TODO Hmmm, I think use Pipline all for serial download
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

func ConstructorMiddleware() *middleware.Middleware {
	if Middleware == nil {
		Middleware = &middleware.Middleware{}
	}
	return Middleware
}

func ConstructorConfig() config.Configurable {
	if App.Config == nil {
		App.Config = json.NewConfig()
		App.Config.Load()
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
	return mutations
}

func startQueries() map[string]*graphql.Field {
	return queries
}
