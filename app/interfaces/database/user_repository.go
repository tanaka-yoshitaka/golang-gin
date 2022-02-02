package database

import (
	"errors"

	"github.com/44taka/golang-gin/domain"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}

	return user, nil
}
