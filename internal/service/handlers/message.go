package handlers

import (
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"github.com/littlebugger/tinode4chat/internal/service/usecase"
	"github.com/littlebugger/tinode4chat/pkg/auth"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// MessageHandler defines the handler for message-related endpoints
type MessageHandler struct {
	uc *usecase.MessageService
}

func NewMessageHandler(service *usecase.MessageService) *MessageHandler {
	return &MessageHandler{uc: service}
}

// SendMessageToChatRoom handles sending messages to a chat room
func (h *MessageHandler) SendMessageToChatRoom(c echo.Context, id string) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	// Parse request body for message details
	var messageRequest struct {
		Content string `json:"content"`
	}
	if err := c.Bind(&messageRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Malformed UserID"})
	}

	// Create a new message entity
	msg := entity.Message{
		Content:    messageRequest.Content,
		Author:     userID,
		ChatRoomID: id,
		Timestamp:  time.Now(),
	}

	// Call the use case to send the message
	msgID, err := h.uc.CreateMessage(ctx, msg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Message sent " + *msgID})
}

// GetChatRoomMessages retrieves messages from a chatroom
func (h *MessageHandler) GetChatRoomMessages(c echo.Context, id string) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Malformed UserID"})
	}

	// Fetch the list of chat rooms from the use case
	chatRooms, err := h.uc.GetMessagesByChatRoom(ctx, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, chatRooms)
}
