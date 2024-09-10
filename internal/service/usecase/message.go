package usecase

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message entity.Message) (primitive.ObjectID, error)
	GetMessagesByChatRoom(ctx context.Context, chatRoomID primitive.ObjectID) ([]entity.Message, error)
}
