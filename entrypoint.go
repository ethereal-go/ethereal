package ethereal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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
	// library gorm for work database
	Db *gorm.DB
	// localization application
	I18n            *i18n.I18n
	Middleware      *Middleware
	GraphQlMutation graphql.Fields
	GraphQlQuery    graphql.Fields
	Context         context.Context
}

func Start() {
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
	}

	App.Middleware.LoadApplication(&App.Context)

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
		//h := handler.New(&handler.Config{
		//	Schema: &schema,
		//	Pretty: true,
		//})
		//ctx := context.WithValue(context.Background(), "test", "get from context")
		//h.ContextHandler(ctx)

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			opts := handler.NewRequestOptions(r)
			result := graphql.Do(graphql.Params{
				Schema:         schema,
				OperationName:  opts.OperationName,
				VariableValues: opts.Variables,
				RequestString:  opts.Query,
				Context:        context.WithValue(context.Background(), "test", "test from context"),
			})
			if len(result.Errors) > 0 {
				log.Printf("wrong result, unexpected errors: %v", result.Errors)
				return
			}
			json.NewEncoder(w).Encode(result)
		})
		//myHandler := http.HandlerFunc(myApp)
		//alice.New(App.Middleware.includeMiddleware...).Then(h)
		// here can add middleware
		http.Handle("/graphql", alice.New(App.Middleware.includeMiddleware...).Then(h))

		http.HandleFunc("/auth0/login", func(w http.ResponseWriter, r *http.Request) {
			claims := EtherealClaims{
				jwt.StandardClaims{
					ExpiresAt: 15000,
					Issuer:    "test",
				},
			}
			// TODO add choose crypt via configuration!
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(JWTKEY())
			fmt.Println(tokenString, err)
		})

		// Serve static files, if variable env debug in true.
		if os.Getenv("DEBAG") != "" && os.Getenv("DEBAG") == "true" {
			_, filename, _, _ := runtime.Caller(0)
			fs := http.FileServer(http.Dir(path.Dir(filename) + "/static"))
			http.Handle("/", fs)
		}

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
