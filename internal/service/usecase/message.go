package usecase

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message entity.Message) (*entity.MessageID, error)
	GetMessagesByChatRoom(ctx context.Context, roomID entity.ChatRoomID) ([]entity.Message, error)
	CheckIfUserInRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) (bool, error)
}

type MessageService struct {
	repo MessageRepository
}

func NewMessageUseCase(repo MessageRepository) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (m *MessageService) CreateMessage(ctx context.Context, message entity.Message) (*entity.MessageID, error) {
	ok, err := m.repo.CheckIfUserInRoom(ctx, message.ChatRoomID, message.Author)
	if err != nil {
		return nil, entity.ErrDbFailed
	}
	if !ok {
		return nil, entity.ErrChatRoomNotFound
	}

	return m.repo.CreateMessage(ctx, message)
}

func (m *MessageService) GetMessagesByChatRoom(ctx context.Context, roomID entity.ChatRoomID, uid entity.UserID) ([]entity.Message, error) {
	ok, err := m.repo.CheckIfUserInRoom(ctx, roomID, uid)
	if err != nil {
		return nil, entity.ErrDbFailed
	}

	if !ok {
		return nil, entity.ErrChatRoomNotFound
	}

	return m.repo.GetMessagesByChatRoom(ctx, roomID)
}
