package memory

import (
	"context"
	"github.com/RafaArauj/study-go/internal/domains/entities"
)

type UserStorage struct {
	user      map[string]*entities.User
	loginToID map[string]string
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		user:      make(map[string]*entities.User),
		loginToID: make(map[string]string),
	}
}

func (u *UserStorage) GetByID(ctx context.Context, id string) (*entities.User, error) {
	user, exists := u.user[id]
	if !exists {
		return nil, entities.ErrUserNotFound
	}

	return user, nil
}

func (u *UserStorage) GetByLogin(ctx context.Context, login string) (*entities.User, error) {
	id, exists := u.loginToID[login]
	if !exists {
		return nil, entities.ErrUserNotFound
	}

	return u.GetByID(ctx, id)
}

func (u *UserStorage) CreateUser(ctx context.Context, user *entities.User) error {
	if _, exists := u.user[user.ID]; exists {
		return entities.ErrUserConflict
	}
	if _, exists := u.loginToID[user.Login]; exists {
		return entities.ErrUserConflict
	}

	u.user[user.ID] = user
	u.loginToID[user.Login] = user.ID
	return nil
}
