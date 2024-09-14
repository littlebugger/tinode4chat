package usecase

import (
	"context"
	"fmt"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message entity.Message) (*entity.MessageID, error)
	GetMessagesByChatRoom(ctx context.Context, roomID entity.ChatRoomID) ([]entity.Message, error)
	CheckIfUserInRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) (bool, error)
}

type MessagingClient interface {
	SendMessage(topicName, messageContent string) error
	Subscribe(topicName string) error
	GetMessages(topicName string) ([]entity.Message, error)
}

type MessageService struct {
	repo            MessageRepository
	messagingClient MessagingClient
}

func NewMessageUseCase(repo MessageRepository, client MessagingClient) *MessageService {
	return &MessageService{
		repo:            repo,
		messagingClient: client,
	}
}

func (m *MessageService) CreateMessage(ctx context.Context, message entity.Message) (*entity.MessageID, error) {
	// Check if the user is in the chat room
	ok, err := m.repo.CheckIfUserInRoom(ctx, message.ChatRoomID, message.Author)
	if err != nil {
		return nil, entity.ErrDbFailed
	}
	if !ok {
		return nil, entity.ErrUserNotInChatRoom
	}

	// Send the message via Tinode
	if err := m.messagingClient.SendMessage(message.ChatRoomID, message.Content); err != nil {
		return nil, fmt.Errorf("failed to send message via Tinode: %w", err)
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

func (m *MessageService) SyncMessagesByChatRoom(ctx context.Context, roomID entity.ChatRoomID, uid entity.UserID) ([]entity.Message, error) {
	// Check if the user is in the chat room
	ok, err := m.repo.CheckIfUserInRoom(ctx, roomID, uid)
	if err != nil {
		return nil, entity.ErrDbFailed
	}
	if !ok {
		return nil, entity.ErrUserNotInChatRoom
	}

	// Subscribe to the topic to receive messages
	if err := m.messagingClient.Subscribe(roomID); err != nil {
		return nil, fmt.Errorf("failed to subscribe to chat room in Tinode: %w", err)
	}

	// Fetch messages from Tinode
	messages, err := m.messagingClient.GetMessages(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages from Tinode: %w", err)
	}

	// Map Tinode messages to your `entity.Message` type
	// ...

	return messages, nil
}
