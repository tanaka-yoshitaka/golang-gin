package database

import (
	"errors"

	"github.com/44taka/golang-gin/domain"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (repo *UserRepository) FindAll(db *gorm.DB) (users []domain.Users, err error) {
	users = []domain.Users{}
	r := db.Find(&users)
	if r.Error != nil {
		return []domain.Users{}, errors.New("user is not found")
	}

	return users, nil
}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}
	r := db.First(&user, id)
	if r.Error != nil {
		return domain.Users{}, errors.New("user is not found")
	}

	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, name string, password string) (user domain.Users, err error) {
	user.Name = name
	user.Password = password
	r := db.Create(&user)
	if r.Error != nil {
		return domain.Users{}, errors.New("failed create user")
	}

	return user, nil
}

func (repo *UserRepository) Update(db *gorm.DB, id int, name string) (user domain.Users, err error) {
	user = domain.Users{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	user.Name = name

	r := db.Save(&user)
	if r.Error != nil {
		return domain.Users{}, errors.New("failed update user")
	}

	return user, nil
}

func (repo *UserRepository) Delete(db *gorm.DB, id int) (user domain.Users, err error) {
	user.ID = id
	r := db.Delete(&user)
	if r.Error != nil {
		return domain.Users{}, errors.New("failed delete user")
	}

	return user, nil
}

func (repo *UserRepository) Login(db *gorm.DB, id int, password string) (user domain.Users, err error) {
	r := db.Where("id = ?", id).Where("password = ?", password).First(&user)
	if r.Error != nil {
		return domain.Users{}, errors.New("user is not found")
	}

	return user, nil
}
