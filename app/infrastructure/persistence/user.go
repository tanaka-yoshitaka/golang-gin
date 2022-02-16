package persistence

import (
	"errors"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.IUserRepository {
	return &UserPersistence{db: db}
}

func (up *UserPersistence) FindAll(ctx *gin.Context) ([]*model.User, error) {
	users := []*model.User{}
	r := up.db.Find(&users)
	if r.Error != nil {
		return users, errors.New("user is not found")
	}
	return users, nil
}

func (up *UserPersistence) FindById(ctx *gin.Context, id int) (model.User, error) {
	user := model.User{}
	r := up.db.First(&user, id)
	if r.Error != nil {
		return user, errors.New("user is not found")
	}
	return user, nil
}

func (up *UserPersistence) Create(ctx *gin.Context, name string, password string) error {
	user := model.User{Name: name, Password: password}
	r := up.db.Create(&user)
	if r.Error != nil {
		return errors.New("failed create user...")
	}
	return nil
}

func (up *UserPersistence) Update(ctx *gin.Context, id int, name string, password string) error {
	user := model.User{}
	up.db.First(&user, id)
	if user.ID <= 0 {
		return errors.New("user is not found")
	}
	user.Name = name
	user.Password = password
	r := up.db.Save(&user)
	if r.Error != nil {
		return errors.New("failed update user...")
	}
	return nil
}

func (up *UserPersistence) Delete(ctx *gin.Context, id int) error {
	user := model.User{ID: id}
	r := up.db.Delete(&user)
	if r.Error != nil {
		return errors.New("failed delete user")
	}
	return nil
}
