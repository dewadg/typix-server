package repositories

import (
	"context"
	"time"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRaceRepository struct {
	db *mongo.Database
}

func NewMongoRaceRepository(db *mongo.Database) RaceRepository {
	return &mongoRaceRepository{
		db: db,
	}
}

func (repo *mongoRaceRepository) Create(ctx context.Context, data models.Race) (models.Race, error) {
	now := time.Now()
	data.ID = primitive.NewObjectID()
	data.CreatedAt = now
	data.UpdatedAt = now

	result, err := repo.db.Collection("races").InsertOne(ctx, data)
	if err != nil {
		return models.Race{}, err
	}
	return repo.Find(ctx, result.InsertedID.(primitive.ObjectID))
}

func (repo *mongoRaceRepository) AddParticipant(ctx context.Context, data models.User) (models.Race, error) {
	panic("implement me")
}

func (repo *mongoRaceRepository) Find(ctx context.Context, id primitive.ObjectID) (models.Race, error) {
	result := repo.db.Collection("races").FindOne(ctx, bson.M{
		"_id":       id,
		"deletedat": nil,
	})
	if result.Err() != nil {
		// TODO: Add proper 404 error
		return models.Race{}, result.Err()
	}

	var race models.Race
	err := result.Decode(&race)

	return race, err
}
