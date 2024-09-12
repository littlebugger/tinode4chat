package repository

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateMessage inserts a new message into a chat room
func (repo *MongoRepository) CreateMessage(ctx context.Context, message entity.Message) (*entity.MessageID, error) {
	collection := repo.GetCollection("messages")

	result, err := collection.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

// GetMessagesByChatRoom retrieves all messages for a specific chat room
func (repo *MongoRepository) GetMessagesByChatRoom(ctx context.Context, chatRoomID entity.ChatRoomID) ([]entity.Message, error) {
	collection := repo.GetCollection("messages")

	var messages []entity.Message
	cursor, err := collection.Find(ctx, bson.M{"chatRoomID": chatRoomID})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
