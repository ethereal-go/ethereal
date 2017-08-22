package ethereal

import "github.com/graphql-go/graphql"

func s()  {
	humanType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A humanoid creature in the Star Wars universe.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"friends": &graphql.Field{
				Type:        graphql.NewList(characterInterface),
				Description: "The friends of the human, or an empty list if they have none.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.Friends, nil
					}
					return []interface{}{}, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.AppearsIn, nil
					}
					return nil, nil
				},
			},
			"homePlanet": &graphql.Field{
				Type:        graphql.String,
				Description: "The home planet of the human, or null if unknown.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.HomePlanet, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: humanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetHuman(p.Args["id"].(int)), nil
				},
			},
		},
	})
}

