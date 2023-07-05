package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() (*mongo.Client, error) {

	const uri = "mongodb+srv://new-user:new-user@cluster0.grve526.mongodb.net/"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	// Ping the MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		return nil, err
	}

	fmt.Println("Connected to MongoDB successfully!")
	return client, nil
}

func DisconnectDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Disconnect(ctx)
}
