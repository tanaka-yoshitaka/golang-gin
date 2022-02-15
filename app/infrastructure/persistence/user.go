package persistence

import (
	"errors"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/domain/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userPersistence struct {
	conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{conn: conn}
}

func (up userPersistence) FindAll(ctx *gin.Context) ([]*model.User, error) {
	users := []*model.User{}
	r := up.conn.Find(&users)
	if r.Error != nil {
		return users, errors.New("user is not found")
	}
	return users, nil
}

func (up userPersistence) FindById(ctx *gin.Context, id int) (model.User, error) {
	user := model.User{}
	r := up.conn.First(&user, id)
	if r.Error != nil {
		return user, errors.New("user is not found")
	}
	return user, nil
}

func (up userPersistence) Create(ctx *gin.Context, name string, password string) error {
	user := model.User{Name: name, Password: password}
	r := up.conn.Create(&user)
	if r.Error != nil {
		return errors.New("failed create user...")
	}
	return nil
}
