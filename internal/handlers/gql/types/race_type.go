package types

import "github.com/graphql-go/graphql"

var RaceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Race",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(ObjectID),
		},
		"words": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
		"participants": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(UserType)),
		},
		"createdAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"updatedAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
