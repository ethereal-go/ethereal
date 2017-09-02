package ethereal

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	//"fmt"
	"github.com/graphql-go/graphql"
)

// Here all constructors application, which return some structure...

func ConstructorI18N() *i18n.I18n {
	//fmt.Println(app)
	if App.I18n == nil {
		App.I18n = i18n.New(
			database.New(ConstructorDb()),
		)
	}
	return App.I18n
}

func ConstructorDb() *gorm.DB {
	if App.Db == nil {
		envLoading()
		App.Db = Database()
	}
	return App.Db

}

func ConstructorMiddleware() *Middleware {
	if App.Middleware == nil {
		App.Middleware = &Middleware{allMiddleware: []AddMiddleware{
			// list standard
			middlewareJWTToken{},
		}}
	}
	return App.Middleware
}

func Mutations() GraphQlMutations {
	return mutations
}

func AddMutations(name string, field *graphql.Field) {
	mutations[name] = field
}

func Queries() GraphQlQueries {
	return queries
}
