package gql

import (
	"os"

	"github.com/dewadg/typix-server/internal"
	"github.com/dewadg/typix-server/internal/handlers/gql/resolvers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewGQLHandler(container *internal.Container) (*handler.Handler, error) {
	queries := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": resolvers.FindUser(container.UserService),
		},
	})

	mutations := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createRace": nil,
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries,
		Mutation: mutations,
	})
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     os.Getenv("APP_ENV") != "production",
		Playground: os.Getenv("APP_ENV") != "production",
	}), nil
}
