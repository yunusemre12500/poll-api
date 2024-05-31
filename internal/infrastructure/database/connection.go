package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	options := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(10).
		SetMinPoolSize(3).
		SetHeartbeatInterval(5 * time.Second)

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		return nil, err
	}

	return client, nil
}
