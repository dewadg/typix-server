package repositories

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Find(ctx context.Context, id primitive.ObjectID) (models.User, error)

	GetByIDs(ctx context.Context, ids []primitive.ObjectID) ([]models.User, error)

	Search(ctx context.Context, keyword string) ([]models.User, error)
}

type RaceRepository interface {
	Create(ctx context.Context, data models.Race) (models.Race, error)

	AddParticipant(ctx context.Context, data models.User) (models.Race, error)

	Find(ctx context.Context, id primitive.ObjectID) (models.Race, error)
}

type WordRepository interface {
	Get(ctx context.Context, count int) ([]string, error)
}
