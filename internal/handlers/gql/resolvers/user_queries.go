package resolvers

import (
	"github.com/dewadg/typix-server/internal/handlers/gql/types"
	"github.com/dewadg/typix-server/internal/services"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUser(userService services.UserService) *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.ObjectID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user, err := userService.Find(p.Context, p.Args["id"].(primitive.ObjectID))
			if err != nil {
				return nil, err
			}
			return user, nil
		},
	}
}
