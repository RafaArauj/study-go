package services

import (
	"context"
	"github.com/RafaArauj/study-go/internal/domains/entities"
	"github.com/RafaArauj/study-go/internal/domains/repositories"
)

type UserService struct {
	storage repositories.UsersStorage
	id      repositories.IDGen
	hash    repositories.HashGen
}

func (u *UserService) CreateUser(ctx context.Context, login, password, fullName string) error {
	hashPwd := u.hash.GenerateFromString(password)
	user := &entities.User{
		ID:       u.id.GenerateID(),
		Login:    login,
		Password: hashPwd,
		FullName: fullName,
	}
	return u.storage.CreateUser(ctx, user)
}

func (u *UserService) ValidateUser(ctx context.Context, login, password string) error {
	hashPwd := u.hash.GenerateFromString(password)
	user, err := u.storage.GetByID(ctx, login)
	if err != nil {
		return err
	}
	if user.Password != hashPwd {
		return entities.ErrInvalidCredentials
	}
	return nil
}
