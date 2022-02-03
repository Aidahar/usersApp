package service

import (
	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/Aidahar/filmsApi/internal/repository/psql"
)

type Servicecer interface {
	GetAllUsers() ([]domain.User, error)
}

type UserService struct {
	repo psql.Userer
}

func NewUsersService(repo psql.Userer) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}
