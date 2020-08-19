package repositories

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRaceRepository struct {
	db *mongo.Database
}

func (repo *mongoRaceRepository) Create(ctx context.Context, data models.Race) (models.Race, error) {
	panic("implement me")
}

func (repo *mongoRaceRepository) AddParticipant(ctx context.Context, data models.User) (models.Race, error) {
	panic("implement me")
}

func (repo *mongoRaceRepository) Find(ctx context.Context, id primitive.ObjectID) (models.Race, error) {
	panic("implement me")
}
