package services

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Find(ctx context.Context, id primitive.ObjectID) (models.User, error)
}
