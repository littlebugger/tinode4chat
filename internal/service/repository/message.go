package repository

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepositoryImpl struct {
	Collection *mongo.Collection
}

func (r *MessageRepositoryImpl) CreateMessage(ctx context.Context, message entity.Message) (primitive.ObjectID, error) {
	result, err := r.Collection.InsertOne(ctx, message)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *MessageRepositoryImpl) GetMessagesByChatRoom(ctx context.Context, chatRoomID primitive.ObjectID) ([]entity.Message, error) {
	var messages []entity.Message
	cursor, err := r.Collection.Find(ctx, bson.M{"chatRoomID": chatRoomID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
