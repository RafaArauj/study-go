package repositories

import (
	"context"
	"github.com/RafaArauj/study-go/internal/domains/entities"
)

type UsersStorage interface {
	GetByID(ctx context.Context, id string) (*entities.User, error)
	GetByLogin(ctx context.Context, login string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
}

type UserService interface {
	CreateUser(ctx context.Context, login, password, fullName string) error
	ValidateUser(ctx context.Context, login, password string) error
}
