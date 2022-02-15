package repository

import (
	"github.com/44taka/golang-gin/domain/model"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	FindAll(ctx *gin.Context) ([]*model.User, error)
	FindById(ctx *gin.Context, id int) (model.User, error)
	Create(ctx *gin.Context, name string, password string) error
}
