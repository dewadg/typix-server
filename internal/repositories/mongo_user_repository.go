package repositories

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) UserRepository {
	return &mongoUserRepository{
		db: db,
	}
}

func (repo *mongoUserRepository) Find(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	result := repo.db.Collection("users").FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		// TODO: Add proper 404 error
		return models.User{}, result.Err()
	}

	var user models.User
	err := result.Decode(&user)

	return user, err
}
