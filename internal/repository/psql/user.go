package psql

import (
	"github.com/Aidahar/filmsApi/ewrap"
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

func (u *UserRepository) AddUser(user domain.User) error {
	u.db.Create(&user)
	return nil
}

func (u *UserRepository) UpdateUser(id int, user domain.User) error {
	if err := u.db.Where("id=?", user.ID).First(&user).Error; err != nil {
		ewrap.LogFatal(err)
	}
	u.db.Save(&user)
	return nil
}

func (u *UserRepository) DeleteUser(id int) error {
	var user domain.User
	u.db.Delete(&user, id)
	return nil
}
