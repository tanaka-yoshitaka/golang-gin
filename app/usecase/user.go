package usecase

import (
	"strconv"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/domain/repository"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	FindAll(*gin.Context) ([]*model.User, error)
	FindById(*gin.Context) (model.User, error)
	Create(ctx *gin.Context) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) FindAll(ctx *gin.Context) (users []*model.User, err error) {
	users, err = uu.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu userUseCase) FindById(ctx *gin.Context) (user model.User, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err = uu.userRepository.FindById(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu userUseCase) Create(ctx *gin.Context) (err error) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	err = uu.userRepository.Create(ctx, name, password)
	if err != nil {
		return err
	}
	return nil
}
