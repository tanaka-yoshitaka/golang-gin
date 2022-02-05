package controllers

import (
	"fmt"
	"strconv"

	"github.com/44taka/golang-gin/interfaces/database"
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

// ユーザー情報取得
func (controller *UsersController) Get(c Context) {
	// 試しにcontextから取得する
	p := c.MustGet("language_message").(*message.Printer)
	fmt.Println(p.Sprintf("bread"))

	id, _ := strconv.Atoi(c.Param("id"))
	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.StatusCode, res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewHandler(res.StatusCode, "OK", user))
}

// ユーザー新規登録
func (controller *UsersController) Post(c Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	_, res := controller.Interactor.Create(id, name)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.StatusCode, res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}

// ユーザー情報更新
func (controller *UsersController) Put(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	_, res := controller.Interactor.Update(id, name)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.StatusCode, res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}

// ユーザー情報削除
func (controller *UsersController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, res := controller.Interactor.Delete(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.StatusCode, res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}

// ログイン
func (controller *UsersController) Login(c Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	password := c.PostForm("password")
	_, token, res := controller.Interactor.Login(id, password)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.StatusCode, res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewHandler(res.StatusCode, "OK", gin.H{"code": token}))
}
