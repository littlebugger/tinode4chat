package usecase

import (
	"context"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (primitive.ObjectID, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
}
