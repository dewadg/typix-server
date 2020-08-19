package repositories

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Find(ctx context.Context, id primitive.ObjectID) (models.User, error)
}

type RaceRepository interface {
	Create(ctx context.Context, data models.Race) (models.Race, error)

	AddParticipant(ctx context.Context, data models.User) (models.Race, error)

	Find(ctx context.Context, id primitive.ObjectID) (models.Race, error)
}
