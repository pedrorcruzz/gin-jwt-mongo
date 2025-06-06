package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryContext struct {
	*mongo.Collection
	client *mongo.Client
}

func NewMongoRepositoryContext(uri, dbName, collectionName string) (*MongoRepositoryContext, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping MongoDB: %v", err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepositoryContext{
		Collection: collection,
		client:     client,
	}, nil
}

func (r *MongoRepositoryContext) Create(contextServer context.Context, document any) error {
	_, err := r.Collection.InsertOne(contextServer, document)
	if err != nil {
		return fmt.Errorf("error inserting document: %v", err)
	}
	return nil
}
