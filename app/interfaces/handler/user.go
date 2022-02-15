package handler

import (
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) FindAll(ctx *gin.Context) {
	// ユースケース呼び出し
	users, err := uh.userUseCase.FindAll(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "not found user",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get user all",
		"result":  users,
	})
	return
}

func (uh userHandler) FindById(ctx *gin.Context) {
	user, err := uh.userUseCase.FindById(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "not found user",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get user",
		"result":  user,
	})
	return
}

func (uh userHandler) Create(ctx *gin.Context) {
	err := uh.userUseCase.Create(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "failed...",
		})
		return
	}
	ctx.AbortWithStatus(201)
	return
}
