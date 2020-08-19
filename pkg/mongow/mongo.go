package mongow

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo(host, port, user, password, authSource, ssl string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=%s&ssl=%s",
		user,
		password,
		host,
		port,
		authSource,
		ssl,
	)
	opts := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
