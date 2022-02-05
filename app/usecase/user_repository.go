package usecase

import (
	"github.com/44taka/golang-gin/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(db *gorm.DB) (users []domain.Users, err error)
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
	Create(db *gorm.DB, id int, name string) (user domain.Users, err error)
	Update(db *gorm.DB, id int, name string) (user domain.Users, err error)
	Delete(db *gorm.DB, id int) (user domain.Users, err error)
	Login(db *gorm.DB, id int, password string) (user domain.Users, err error)
}
