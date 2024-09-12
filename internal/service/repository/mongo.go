package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// ConnectToMongo establishes a connection to MongoDB and returns a MongoRepository
func ConnectToMongo(ctx context.Context, uri, dbName string) (*MongoRepository, error) {
	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check the connection by pinging MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")

	// Return the MongoRepository with the connected client and database
	return &MongoRepository{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

// CloseMongoConnection closes the MongoDB connection gracefully
func (repo *MongoRepository) CloseMongoConnection(ctx context.Context) error {
	if err := repo.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("error closing MongoDB connection: %v", err)
	}

	log.Println("MongoDB connection closed.")
	return nil
}

// GetCollection returns a MongoDB collection reference
func (repo *MongoRepository) GetCollection(collectionName string) *mongo.Collection {
	return repo.Database.Collection(collectionName)
}
