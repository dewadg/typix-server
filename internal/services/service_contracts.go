package services

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Find(ctx context.Context, id primitive.ObjectID) (models.User, error)

	Search(ctx context.Context, keyword string) ([]models.User, error)
}

type RaceService interface {
	Create(ctx context.Context, data models.Race) (models.Race, error)

	Find(ctx context.Context, id primitive.ObjectID) (models.Race, error)
}
