package sqlit

import (
	"github.com/Aidahar/filmsApi/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type Userer interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	AddUser(user domain.User) error
	UpdateUser(id int, user domain.User) error
	DeleteUser(id int) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
