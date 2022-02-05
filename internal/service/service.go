package service

import (
	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/Aidahar/filmsApi/internal/repository/psql"
)

type Servicecer interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	AddUser(user domain.User) (int, error)
	UpdateUser(id int, user domain.User) error
	DeleteUser(id int) error
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

func (s *UserService) GetUserById(id int) (domain.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) AddUser(user domain.User) error {
	return s.repo.AddUser(user)
}

func (s *UserService) UpdateUser(id int, user domain.User) error {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
