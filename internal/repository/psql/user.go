package psql

import (
	"github.com/Aidahar/filmsApi/internal/domain"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	u.db.Find(&users)
	return users, nil
}
