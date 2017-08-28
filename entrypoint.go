package ethereal

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"net/http"
	"path"
	"runtime"
	"strconv"
)

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"display_name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var usersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"role": &graphql.Field{
			Type: roleType,
		},
	},
})

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
		"users": &graphql.Field{
			Type:        graphql.NewList(usersType),
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var users []User
				var role Role
				app.Db.Find(&users)

				idQuery, isOK := params.Args["id"].(string)

				if isOK {
					for _, user := range users {
						if strconv.Itoa(int(user.ID)) == idQuery {
							app.Db.Model(&user).Related(&role)
							user.Role = role
							fmt.Println(user.Role)
							return []User{user}, nil
						}
					}
				}
				return users, nil
			},
		},
		"role": &graphql.Field{
			Type:        graphql.NewList(roleType),
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var roles []Role
				app.Db.Find(&roles)

				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					for _, role := range roles {
						if string(role.ID) == idQuery {
							return role, nil
						}
					}
				}

				return roles, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	//Mutation: rootMutation,
})

func Start() {
	envLoading()
	db := Database()
	I18n := i18n.New(
		database.New(db),
	)
	app = App{Db: Database(), I18n: I18n}
	//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	//	result := executeQuery(r.URL.Query().Get("query"), schema)
	//	json.NewEncoder(w).Encode(result)
	//})
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	// Serve static files
	_, filename, _, _ := runtime.Caller(0)
	//fmt.Println(path.Dir(filename))
	fs := http.FileServer(http.Dir(path.Dir(filename) + "/static"))
	http.Handle("/", fs)
	fmt.Println("Now server is running on port 8080")

	//fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
	//fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")

	http.ListenAndServe(":8080", nil)
}
