package usecase

import (
	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/domain/repository"
	"github.com/gin-gonic/gin"
)

type IUserUseCase interface {
	FindAll(ctx *gin.Context) ([]*model.User, error)
	FindById(ctx *gin.Context, id int) (model.User, error)
	Create(ctx *gin.Context, name string, password string) error
	Update(ctx *gin.Context, id int, name string, password string) error
	Delete(ctx *gin.Context, id int) error
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

func (uu *UserUseCase) FindAll(ctx *gin.Context) (users []*model.User, err error) {
	users, err = uu.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *UserUseCase) FindById(ctx *gin.Context, id int) (user model.User, err error) {
	user, err = uu.userRepository.FindById(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu *UserUseCase) Create(ctx *gin.Context, name string, password string) (err error) {
	err = uu.userRepository.Create(ctx, name, password)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUseCase) Update(ctx *gin.Context, id int, name string, password string) (err error) {
	err = uu.userRepository.Update(ctx, id, name, password)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUseCase) Delete(ctx *gin.Context, id int) (err error) {
	err = uu.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
