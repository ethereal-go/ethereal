package ethereal

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"net/http"
	"os"
	"path"
	"runtime"
)

var app App

type App struct {
	Db   *gorm.DB
	I18n *i18n.I18n
	//Locale string // type localization, example(ru-RU, en-US)
}

// root mutation
//var rootMutation = graphql.NewObject(graphql.ObjectConfig{
//	Name: "RootMutation",
//	Fields: graphql.Fields{
//		/*
//			curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
//		*/
//		"createTodo": &graphql.Field{
//			Type:        todoType, // the return type for this field
//			Description: "Create new todo",
//			Args: graphql.FieldConfigArgument{
//				"text": &graphql.ArgumentConfig{
//					Type: graphql.NewNonNull(graphql.String),
//				},
//			},
//			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//
//				// marshall and cast the argument value
//				text, _ := params.Args["text"].(string)
//
//				// figure out new id
//				newID := RandStringRunes(8)
//
//				// perform mutation operation here
//				// for e.g. create a Todo and save to DB.
//				newTodo := Todo{
//					ID:   newID,
//					Text: text,
//					Done: false,
//				}
//
//				TodoList = append(TodoList, newTodo)
//
//				// return the new Todo object that we supposedly save to DB
//				// Note here that
//				// - we are returning a `Todo` struct instance here
//				// - we previously specified the return Type to be `todoType`
//				// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
//				return newTodo, nil
//			},
//		},
//		/*
//			curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
//		*/
//		"updateTodo": &graphql.Field{
//			Type:        todoType, // the return type for this field
//			Description: "Update existing todo, mark it done or not done",
//			Args: graphql.FieldConfigArgument{
//				"done": &graphql.ArgumentConfig{
//					Type: graphql.Boolean,
//				},
//				"id": &graphql.ArgumentConfig{
//					Type: graphql.NewNonNull(graphql.String),
//				},
//			},
//			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//				// marshall and cast the argument value
//				done, _ := params.Args["done"].(bool)
//				id, _ := params.Args["id"].(string)
//				affectedTodo := Todo{}
//
//				// Search list for todo with id and change the done variable
//				for i := 0; i < len(TodoList); i++ {
//					if TodoList[i].ID == id {
//						TodoList[i].Done = done
//						// Assign updated todo so we can return it
//						affectedTodo = TodoList[i]
//						break
//					}
//				}
//				// Return affected todo
//				return affectedTodo, nil
//			},
//		},
//	},
//})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &UserField,
		"role":  &RoleField,
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	//Mutation: rootMutation,
})

func ConstructorDb() *gorm.DB {
	if app.Db == nil {
		envLoading()
		app.Db = Database()
	}
	return app.Db

}
func ConstructorI18N() *i18n.I18n {
	if app.I18n == nil {
		app.I18n = i18n.New(
			database.New(ConstructorDb()),
		)
	}
	return app.I18n
}

func Start() {
	//envLoading()
	//db := Database()
	//I18n := i18n.New(
	//	database.New(db),
	//)
	fmt.Println(string(ConstructorI18N().Scope("graphQL").T(os.Getenv("LOCALE"), "User.Description")))

	app = App{Db: ConstructorDb(), I18n: ConstructorI18N()}
	I18nGraphQL().Fill()
	if len(os.Args) > 1 {
		CliRun()
	} else {
		//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		//	result := executeQuery(r.URL.Query().Get("query"), schema)
		//	json.NewEncoder(w).Encode(result)
		//})
		h := handler.New(&handler.Config{
			Schema: &schema,
			Pretty: true,
		})
		http.Handle("/graphql", middlewareLocal(h))
		// Serve static files
		_, filename, _, _ := runtime.Caller(0)
		fs := http.FileServer(http.Dir(path.Dir(filename) + "/static"))
		http.Handle("/", fs)
		fmt.Println("Now server is running on port 8080")

		//fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
		//fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")

		http.ListenAndServe(":8080", nil)
	}

}
