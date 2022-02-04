package psql

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
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	u.db.Find(&users)
	return users, nil
}

func (u *UserRepository) GetUserById(id int) (domain.User, error) {
	var user domain.User
	u.db.First(&user, id)
	return user, nil
}
