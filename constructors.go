package ethereal

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
)

// Here all constructors application, which return some structure...

func ConstructorI18N() *i18n.I18n {
	if app.I18n == nil {
		app.I18n = i18n.New(
			database.New(ConstructorDb()),
		)
	}
	return app.I18n
}

func ConstructorDb() *gorm.DB {
	if app.Db == nil {
		envLoading()
		app.Db = Database()
	}
	return app.Db

}

func ConstructorMiddleware() *Middleware {
	if app.Middleware == nil {
		app.Middleware = &Middleware{allMiddleware: []AddMiddleware{
			// list standard
			middlewareJWTToken{},
		}}
	}
	return app.Middleware
}

func ConstructorApp() *App {
	if app == nil {
		app = &App{
			Db:         ConstructorDb(),
			I18n:       ConstructorI18N(),
			Middleware: ConstructorMiddleware(),
			GraphQlQuery: graphql.Fields{
				"users": &UserField,
				"role":  &RoleField,
			},
			GraphQlMutation: graphql.Fields{
				"createUser": &createUser,
			},
		}
	}
	return app
}
