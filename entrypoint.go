package ethereal

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

var app App

// Base structure
type App struct {
	Db   *gorm.DB
	I18n *i18n.I18n
}

//root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &createUser,
	},
})

// root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &UserField,
		"role":  &RoleField,
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// yet
func ConstructorI18N() *i18n.I18n {
	if app.I18n == nil {
		app.I18n = i18n.New(
			database.New(ConstructorDb()),
		)
	}
	return app.I18n
}

func Start() {
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
		// here can add middleware
		http.Handle("/graphql", h)

		http.HandleFunc("get-token", func(w http.ResponseWriter, r *http.Request){

		})
		// Serve static files
		_, filename, _, _ := runtime.Caller(0)
		fs := http.FileServer(http.Dir(path.Dir(filename) + "/static"))

		http.Handle("/", fs)

		if os.Getenv("SERVER_PORT") == "" {
			os.Setenv("SERVER_PORT", "8080")
		}
		fmt.Println("Now server is running on port " + os.Getenv("SERVER_PORT"))
		http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	}
}

/**
/ Load environment variables
*/
func envLoading() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
