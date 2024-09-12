package handlers

import (
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"github.com/littlebugger/tinode4chat/internal/service/usecase"
	"github.com/littlebugger/tinode4chat/pkg/auth"
	api "github.com/littlebugger/tinode4chat/pkg/server/chatroom"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ChatRoomHandler defines the handler for chatroom-related endpoints
type ChatRoomHandler struct {
	uc *usecase.ChatRoomUseCase
}

func NewChatRoomHandler(service *usecase.ChatRoomUseCase) *ChatRoomHandler {
	return &ChatRoomHandler{uc: service}
}

// CreateChatRoom handles chatroom creation
func (h *ChatRoomHandler) CreateChatRoom(c echo.Context) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	var roomRequest api.ChatRoomCreate
	if err := c.Bind(&roomRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Create new ChatRoom entity
	chatRoom := entity.ChatRoom{
		Name: roomRequest.Name,
	}

	// Call the use case to create the chat room
	chatID, err := h.uc.CreateChatRoom(ctx, chatRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create chat room"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Chat room created", "roomID": *chatID})
}

// ListChatRooms lists all chatrooms
func (h *ChatRoomHandler) ListChatRooms(c echo.Context) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	// Fetch the list of chat rooms from the use case
	chatRooms, err := h.uc.ListChatRooms(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve chat rooms"})
	}

	return c.JSON(http.StatusOK, chatRooms)
}

// JoinChatRoom allows a user to join a chatroom
func (h *ChatRoomHandler) JoinChatRoom(c echo.Context, id string) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	// Extract roomID and userID from the URL and the token
	roomID := c.Param("id")
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Malformed UserID"})
	}

	// Call the use case to add the user to the chat room
	if err := h.uc.AddUserToChatRoom(ctx, roomID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to join chat room"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Joined chat room"})
}

// LeaveChatRoom allows a user to leave a chatroom
func (h *ChatRoomHandler) LeaveChatRoom(c echo.Context, id string) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	// Extract roomID and userID from the URL and the token
	roomID := c.Param("id")
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Malformed UserID"})
	}

	// Call the use case to remove the user from the chat room
	if err := h.uc.RemoveUserFromChatRoom(ctx, roomID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to left chat room"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Left chat room"})
}
