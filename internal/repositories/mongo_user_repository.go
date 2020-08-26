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

func (repo *mongoUserRepository) GetByIDs(ctx context.Context, ids []primitive.ObjectID) ([]models.User, error) {
	cursor, err := repo.db.Collection("users").Find(ctx, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
		"deletedat": nil,
	})
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = cursor.All(ctx, &users)

	return users, err
}

func (repo *mongoUserRepository) Search(ctx context.Context, keyword string) ([]models.User, error) {
	searchFilter := bson.M{
		"deletedAt": nil,
		"username": bson.M{
			"$regex":   keyword,
			"$options": "i",
		},
	}

	cursor, err := repo.db.Collection("users").Find(ctx, searchFilter)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)
	err = cursor.All(ctx, &users)

	return users, err
}
