package resolvers

import (
	"github.com/dewadg/typix-server/internal/handlers/gql/types"
	"github.com/dewadg/typix-server/internal/models"
	"github.com/dewadg/typix-server/internal/services"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRace(raceService services.RaceService) *graphql.Field {
	return &graphql.Field{
		Type: types.RaceType,
		Args: graphql.FieldConfigArgument{
			"payload": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.CreateRace),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var request = struct {
				ParticipantIDs []primitive.ObjectID `mapstructure:"participantIds"`
			}{}
			mapstructure.Decode(p.Args["payload"], &request)

			data := models.Race{
				ParticipantIDs: request.ParticipantIDs,
			}

			race, err := raceService.Create(p.Context, data)
			if err != nil {
				return nil, err
			}
			return race, nil
		},
	}
}
