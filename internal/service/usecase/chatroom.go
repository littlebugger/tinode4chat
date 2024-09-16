package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/littlebugger/tinode4chat/internal/service/entity"
)

// ChatRoomRepository defines the interface for repository methods
type ChatRoomRepository interface {
	CreateChatRoom(ctx context.Context, room entity.ChatRoom) (*string, error)
	ListChatRooms(ctx context.Context) ([]entity.ChatRoom, error)
	AddUserToChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error
	RemoveUserFromChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error
	IsRoomExist(ctx context.Context, name string) (bool, error)

	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserEmailByID(ctx context.Context, userID entity.UserID) (string, error)
}

type RoomsClient interface {
	CreateTopic(topicName string) (string, error)
	AddUserToTopic(topicName, userEmail string) error
	RemoveUserFromTopic(topicName, userEmail string) error
}

// ChatRoomService defines the business logic for chat room operations
type ChatRoomService struct {
	repo  ChatRoomRepository
	rooms RoomsClient
}

func NewChatRoomUseCase(repo ChatRoomRepository, client RoomsClient) *ChatRoomService {
	return &ChatRoomService{
		repo:  repo,
		rooms: client,
	}
}

// TODO: maybe need to check for existence of user in db not only for signed jwt
// TODO: check if room exist before any action with it

// CreateChatRoom handles the business logic for creating a chat room
func (uc *ChatRoomService) CreateChatRoom(ctx context.Context, room entity.ChatRoom) (*string, error) {
	if room.Name == "" {
		return nil, entity.ErrInvalidRoomName
	}

	// Check if the room already exists in your database
	exists, err := uc.repo.IsRoomExist(ctx, room.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, entity.ErrChatRoomAlreadyExists
	}

	// Create the chat room in Tinode
	topicName, err := uc.rooms.CreateTopic(room.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat room in Tinode: %w", err)
	}

	// Save the chat room in your database
	room.ID = topicName // Assuming topicName is the unique identifier
	return uc.repo.CreateChatRoom(ctx, room)
}

// TODO: wrap all errors into entity errors

// ListChatRooms returns a list of chat rooms
func (uc *ChatRoomService) ListChatRooms(ctx context.Context) ([]entity.ChatRoom, error) {
	return uc.repo.ListChatRooms(ctx)
}

// AddUserToChatRoom adds a user to a chat room
func (uc *ChatRoomService) AddUserToChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	// Get the user email from the user ID
	userEmail, err := uc.repo.GetUserEmailByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user email: %w", err)
	}

	// Add the user to the topic in Tinode
	if err := uc.rooms.AddUserToTopic(roomID, userEmail); err != nil {
		return fmt.Errorf("failed to add user to chat room in Tinode: %w", err)
	}

	// Update your database if necessary
	return uc.repo.AddUserToChatRoom(ctx, roomID, userID)
}

// RemoveUserFromChatRoom removes a user from a chat room
func (uc *ChatRoomService) RemoveUserFromChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	// Get the user email from the user ID
	userEmail, err := uc.repo.GetUserEmailByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user email: %w", err)
	}

	// Remove the user from the topic in Tinode
	if err := uc.rooms.RemoveUserFromTopic(roomID, userEmail); err != nil {
		return fmt.Errorf("failed to remove user from chat room in Tinode: %w", err)
	}

	// Update your database if necessary
	return uc.repo.RemoveUserFromChatRoom(ctx, roomID, userID)
}

func (uc *ChatRoomService) HandleMetaEvent(ctx context.Context, meta map[string]interface{}) error {
	if contacts, ok := meta["sub"].([]interface{}); ok {
		for _, c := range contacts {
			contact := c.(map[string]interface{})
			topic := contact["topic"].(string)
			userID := contact["user"].(string)

			log.Printf("New subscription to topic: %s", topic)

			// Check if the room already exists
			exists, err := uc.repo.IsRoomExist(ctx, topic)
			if err != nil {
				return err
			}

			if !exists {
				room := entity.ChatRoom{
					ID:           topic,
					Name:         topic,
					Participants: []entity.UserID{userID},
				}
				_, err := uc.repo.CreateChatRoom(ctx, room)
				if err != nil {
					return err
				}
			} else {
				err := uc.repo.AddUserToChatRoom(ctx, topic, userID)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
