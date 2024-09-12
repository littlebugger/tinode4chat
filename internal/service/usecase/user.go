package usecase

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/littlebugger/tinode4chat/pkg/auth"
	"golang.org/x/crypto/bcrypt"

	"github.com/littlebugger/tinode4chat/internal/service/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.UserID, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser registers a new user, ensuring the user data is valid and the email is unique. It also hashes the password.
func (uc *UserService) CreateUser(ctx context.Context, user entity.User) (*entity.UserID, error) {
	if !user.Valid() {
		return nil, entity.ErrInvalidUserEntry
	}

	ext, err := uc.GetUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, entity.ErrUserNotFound) {
		return nil, err
	}
	if ext != nil {
		return nil, entity.ErrUserAlreadyExists
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return nil, entity.ErrCryptoFailed
	}

	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return uc.repo.GetUserByEmail(ctx, email)
}

// Login authenticates a user using provided email and password, returning a JWT token if successful.
func (uc *UserService) Login(ctx context.Context, email, password string) (*string, error) {
	// TODO: email need sanitization
	if email == "" || password == "" {
		return nil, entity.ErrInvalidCredentials
	}

	// Find user by email in the repository
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Infof("login failed: %v", err)

		return nil, entity.ErrUserNotFound
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, entity.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := auth.GenerateJWTToken(user)
	if err != nil {
		return nil, entity.ErrCryptoFailed
	}

	return &token, nil
}

// hashPassword hashes a given password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
