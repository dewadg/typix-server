package resolvers

import (
	"github.com/dewadg/typix-server/internal/handlers/gql/types"
	"github.com/dewadg/typix-server/internal/services"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindRace(raceService services.RaceService) *graphql.Field {
	return &graphql.Field{
		Type: types.RaceType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.ObjectID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			race, err := raceService.Find(p.Context, p.Args["id"].(primitive.ObjectID))
			if err != nil {
				return nil, err
			}
			return race, nil
		},
	}
}
