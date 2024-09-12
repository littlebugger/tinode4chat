package handlers

import (
	"errors"
	"github.com/littlebugger/tinode4chat/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"github.com/littlebugger/tinode4chat/internal/service/usecase"
	api "github.com/littlebugger/tinode4chat/pkg/server/user"
)

// UserHandler defines the handler for user-related endpoints
type UserHandler struct {
	uc *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{uc: service}
}

// SignupUser handles user registration
func (h *UserHandler) SignupUser(c echo.Context) error {
	ctx := c.Request().Context()

	// Parse the incoming JSON request
	var userReq api.UserSignup
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user := entity.User{
		Email:    userReq.Email,
		Username: userReq.Username,
		Password: userReq.Password,
	}

	id, err := h.uc.CreateUser(ctx, user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)

		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully", "id": *id})
}

// LoginUser handles user login
func (h *UserHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()

	var userReq api.UserLogin
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	token, err := h.uc.Login(ctx, userReq.Email, userReq.Password)
	if err != nil {
		log.Printf("Failed to login user: %v", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": *token})
}

// GetUserProfile handles retrieving user profile
func (h *UserHandler) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	err := auth.JWTMiddleware(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token missing or invalid"})
	}

	email, ok := c.Get("email").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JWT token malformed or invalid"})
	}

	user, err := h.uc.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	resp := api.UserProfile{
		Email:    &user.Email,
		Username: &user.Username,
	}

	return c.JSON(http.StatusOK, map[string]api.UserProfile{"user": resp})
}
