package handler

import (
	"net/http"

	"github.com/44taka/golang-gin/interfaces/validator"
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserHandler struct {
	userUseCase usecase.IUserUseCase
}

func NewUserHandler(uu usecase.IUserUseCase) IUserHandler {
	return &UserHandler{userUseCase: uu}
}

func (uh *UserHandler) FindAll(ctx *gin.Context) {
	// ユースケース呼び出し
	users, err := uh.userUseCase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user all",
		"result":  users,
	})
	return
}

func (uh *UserHandler) FindById(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
		})
		return
	}
	user, err := uh.userUseCase.FindById(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user",
		"result":  user,
	})
	return
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	var form validator.UserForm
	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
		})
		return
	}
	err = uh.userUseCase.Create(ctx, form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
		})
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
	return
}

func (uh *UserHandler) Update(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
		})
		return
	}
	var form validator.UserForm
	err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
		})
		return
	}
	err = uh.userUseCase.Update(ctx, uri.ID, form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
		})
		return
	}
	ctx.AbortWithStatus(http.StatusNoContent)
	return
}

func (uh *UserHandler) Delete(ctx *gin.Context) {
	var uri validator.UserUri
	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "バリデーションエラー",
		})
		return
	}
	err = uh.userUseCase.Delete(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed...",
		})
		return
	}
	ctx.AbortWithStatus(http.StatusNoContent)
	return
}
