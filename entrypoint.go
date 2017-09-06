package ethereal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/justinas/alice"
	"github.com/qor/i18n"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

var App Application

// Base structure
type Application struct {
	Db              *gorm.DB
	I18n            *i18n.I18n
	Middleware      *Middleware
	GraphQlMutation graphql.Fields
	GraphQlQuery    graphql.Fields
	Context         context.Context
	Config          *Config
}

func Start() {
	// Config variables
	var (
		debug string = config("GRAPHQL.DEBUG").(string)
		host  string = config("HOST.PORT").(string)
	)

	// First we have to determine the mode of operation
	// - cli console
	// - api server
	// Secondly, we must determine the sequence of actions

	App = Application{
		Db:              ConstructorDb(),
		I18n:            ConstructorI18N(),
		Middleware:      ConstructorMiddleware(),
		GraphQlQuery:    startQueries(),
		GraphQlMutation: startMutations(),
		Context:         context.Background(),
		Config:          ConstructorConfig(),
	}

	App.Middleware.LoadApplication()

	//root mutation
	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: App.GraphQlMutation,
	})

	// root query
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: App.GraphQlQuery,
	})

	// define schema, with our rootQuery and rootMutation
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	I18nGraphQL().Fill()

	if len(os.Args) > 1 {
		CliRun()
	} else {

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			opts := handler.NewRequestOptions(r)
			fmt.Println(opts.OperationName)
			result := graphql.Do(graphql.Params{
				Schema:         schema,
				OperationName:  opts.OperationName,
				VariableValues: opts.Variables,
				RequestString:  opts.Query,
				Context:        App.Context,
			})
			if len(result.Errors) > 0 {
				log.Printf("wrong result, unexpected errors: %v", result.Errors)
				return
			}
			json.NewEncoder(w).Encode(result)
		})

		// here can add middleware
		http.Handle("/graphql", alice.New(App.Middleware.includeMiddleware...).Then(h))

		//http.HandleFunc("/auth0/login", func(w http.ResponseWriter, r *http.Request) {
		//	claims := EtherealClaims{
		//		jwt.StandardClaims{
		//			ExpiresAt: 15000,
		//			Issuer:    "test",
		//		},
		//	}
		//	// TODO add choose crypt via configuration!
		//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//
		//	tokenString, err := token.SignedString(JWTKEY())
		//	fmt.Println(tokenString, err)
		//})

		// Serve static files, if variable env debug in true.
		if debug == "true" {
			_, filename, _, _ := runtime.Caller(0)
			fs := http.FileServer(http.Dir(path.Dir(filename) + "/static"))
			http.Handle("/", fs)
		}

		fmt.Println("Now server is running on port " + host)
		http.ListenAndServe(":" + host, nil)
	}
}
