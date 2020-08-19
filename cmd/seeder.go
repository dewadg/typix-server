package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dewadg/typix-server/internal/models"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Seed(client *mongo.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "db:seed",
		Short: "Run DB seeders",
		Run: func(cmd *cobra.Command, args []string) {
			seeders := []func(database *mongo.Database) error{
				seedUsers,
			}

			db := client.Database(os.Getenv("MONGO_DATABASE"))
			for _, seeder := range seeders {
				if err := seeder(db); err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		},
	}
}

func seedUsers(db *mongo.Database) error {
	now := time.Now()

	users := []models.User{
		{
			ID:        primitive.NewObjectID(),
			Username:  "tester1",
			Email:     "tester1@dewadg.id",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        primitive.NewObjectID(),
			Username:  "tester2",
			Email:     "tester2@dewadg.id",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	ctx := context.TODO()
	coll := db.Collection("users")

	for _, user := range users {
		result := coll.FindOne(ctx, bson.M{"username": user.Username})
		if !errors.Is(result.Err(), mongo.ErrNoDocuments) {
			continue
		}

		_, err := coll.InsertOne(ctx, user)
		if err != nil {
			return err
		}
	}
	return nil
}
