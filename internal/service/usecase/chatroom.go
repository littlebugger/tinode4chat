package usecase

import (
	"context"

	"github.com/littlebugger/tinode4chat/internal/service/entity"
)

// ChatRoomRepository defines the interface for repository methods
type ChatRoomRepository interface {
	CreateChatRoom(ctx context.Context, room entity.ChatRoom) (*string, error)
	ListChatRooms(ctx context.Context) ([]entity.ChatRoom, error)
	AddUserToChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error
	RemoveUserFromChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error
	IsRoomExist(ctx context.Context, name string) (bool, error)
}

// ChatRoomUseCase defines the business logic for chat room operations
type ChatRoomUseCase struct {
	repo ChatRoomRepository
}

func NewChatRoomUseCase(repo ChatRoomRepository) *ChatRoomUseCase {
	return &ChatRoomUseCase{
		repo: repo,
	}
}

// TODO: maybe need to check for existence of user in db not only for signed jwt
// TODO: check if room exist before any action with it

// CreateChatRoom handles the business logic for creating a chat room
func (uc *ChatRoomUseCase) CreateChatRoom(ctx context.Context, room entity.ChatRoom) (*string, error) {
	if room.Name == "" {
		return nil, entity.ErrChatRoomNotFound
	}

	return uc.repo.CreateChatRoom(ctx, room)
}

// TODO: wrap all errors into entity errors

// ListChatRooms returns a list of chat rooms
func (uc *ChatRoomUseCase) ListChatRooms(ctx context.Context) ([]entity.ChatRoom, error) {
	return uc.repo.ListChatRooms(ctx)
}

// AddUserToChatRoom adds a user to a chat room
func (uc *ChatRoomUseCase) AddUserToChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	return uc.repo.AddUserToChatRoom(ctx, roomID, userID)
}

// RemoveUserFromChatRoom removes a user from a chat room
func (uc *ChatRoomUseCase) RemoveUserFromChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	return uc.repo.RemoveUserFromChatRoom(ctx, roomID, userID)
}
