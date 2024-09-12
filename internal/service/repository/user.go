package repository

import (
	"context"
	"errors"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser inserts a new user into the "users" collection
func (repo *MongoRepository) CreateUser(ctx context.Context, user entity.User) (*entity.UserID, error) {
	collection := repo.GetCollection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

// GetUserByEmail fetches a user by their email
func (repo *MongoRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	collection := repo.GetCollection("users")

	var user entity.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID fetches a user by their ID
func (repo *MongoRepository) GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error) {
	collection := repo.GetCollection("users")

	var user entity.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user document
func (repo *MongoRepository) UpdateUser(ctx context.Context, user entity.User) error {
	collection := repo.GetCollection("users")

	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (repo *MongoRepository) CheckIfUserInRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) (bool, error) {
	collection := repo.GetCollection("users")

	rid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return false, err
	}

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}

	// Create a filter to find the user by their ID and check if the roomID is in the joinedRooms array
	filter := bson.M{
		"_id":         uid, // Match the user by ID
		"joinedRooms": rid, // Check if the roomID is in the joinedRooms array
	}

	// Try to find a matching document
	err = collection.FindOne(ctx, filter).Err()
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
