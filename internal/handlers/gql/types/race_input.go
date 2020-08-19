package types

import "github.com/graphql-go/graphql"

var CreateRace = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateRace",
	Fields: graphql.InputObjectConfigFieldMap{
		"participantIds": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ObjectID),
		},
	},
})
