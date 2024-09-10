package usecase

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoomRepository interface {
	CreateChatRoom(ctx context.Context, room entity.ChatRoom) (primitive.ObjectID, error)
	GetChatRoomByID(ctx context.Context, id primitive.ObjectID) (*entity.ChatRoom, error)
	AddUserToChatRoom(ctx context.Context, roomID primitive.ObjectID, userID primitive.ObjectID) error
	RemoveUserFromChatRoom(ctx context.Context, roomID primitive.ObjectID, userID primitive.ObjectID) error
}
