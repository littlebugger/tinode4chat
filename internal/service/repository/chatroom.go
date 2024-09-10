package repository

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRoomRepositoryImpl struct {
	Collection *mongo.Collection
}

func (r *ChatRoomRepositoryImpl) CreateChatRoom(ctx context.Context, room entity.ChatRoom) (primitive.ObjectID, error) {
	result, err := r.Collection.InsertOne(ctx, room)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *ChatRoomRepositoryImpl) GetChatRoomByID(ctx context.Context, id primitive.ObjectID) (*entity.ChatRoom, error) {
	var room entity.ChatRoom
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *ChatRoomRepositoryImpl) AddUserToChatRoom(ctx context.Context, roomID primitive.ObjectID, userID primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": roomID}, bson.M{"$addToSet": bson.M{"participants": userID}})
	return err
}

func (r *ChatRoomRepositoryImpl) RemoveUserFromChatRoom(ctx context.Context, roomID primitive.ObjectID, userID primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": roomID}, bson.M{"$pull": bson.M{"participants": userID}})
	return err
}
