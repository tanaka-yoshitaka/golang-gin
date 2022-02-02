package usecase

import (
	"github.com/44taka/golang-gin/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
}
