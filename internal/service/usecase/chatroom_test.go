package usecase

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"github.com/littlebugger/tinode4chat/internal/service/usecase/mocks"
	"testing"
)

func TestChatRoomUseCase_CreateChatRoom(t *testing.T) {
	mockRepo := mocks.NewMockChatRoomRepository(t)
	uc := NewChatRoomUseCase(mockRepo)

	tests := []struct {
		name    string
		room    entity.ChatRoom
		wantErr bool
	}{
		{
			name:    "Positive Case",
			room:    entity.ChatRoom{Name: "Public Room"},
			wantErr: false,
		},
		{
			name:    "Negative Case With Empty Room Name",
			room:    entity.ChatRoom{Name: ""},
			wantErr: true,
		},
		{
			name:    "Negative Case With Existing Room Name",
			room:    entity.ChatRoom{Name: "Public Room"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := uc.CreateChatRoom(context.Background(), tt.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatRoomService.CreateChatRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatRoomUseCase_ListChatRooms(t *testing.T) {
	mockRepo := mocks.NewMockChatRoomRepository(t)
	uc := NewChatRoomUseCase(mockRepo)

	mockRepo.CreateChatRoom(context.Background(), entity.ChatRoom{Name: "Room 1"})
	mockRepo.CreateChatRoom(context.Background(), entity.ChatRoom{Name: "Room 2"})
	mockRepo.CreateChatRoom(context.Background(), entity.ChatRoom{Name: "Room 3"})

	tests := []struct {
		name          string
		wantChatRooms []entity.ChatRoom
		wantErr       bool
	}{
		{
			name:          "Positive Case, three rooms",
			wantChatRooms: []entity.ChatRoom{{Name: "Room 1"}, {Name: "Room 2"}, {Name: "Room 3"}},
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRooms, err := uc.ListChatRooms(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatRoomService.ListChatRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotRooms) != len(tt.wantChatRooms) {
				t.Errorf("ChatRoomService.ListChatRooms() got = %v rooms, want %v rooms", len(gotRooms), len(tt.wantChatRooms))
				return
			}
			for i, gotRoom := range gotRooms {
				if gotRoom.Name != tt.wantChatRooms[i].Name {
					t.Errorf("ChatRoomService.ListChatRooms() got = %v, want %v", gotRoom, tt.wantChatRooms[i])
				}
			}
		})
	}
}
