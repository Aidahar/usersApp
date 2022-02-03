package service

import (
	"context"

	"github.com/Aidahar/filmsApi/internal/domain"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type Users struct {
	repo UsersRepository
}

func NewUsers(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) GetAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAllUsers(ctx)
}
