package controllers

import (
	"fmt"
	"strconv"

	"github.com/44taka/golang-gin/interfaces/database"
	"github.com/44taka/golang-gin/usecase"
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

func (controller *UsersController) Get(c Context) {
	// 試しにcontextから取得する
	p := c.MustGet("language_message").(*message.Printer)
	fmt.Println(p.Sprintf("bread"))

	id, _ := strconv.Atoi(c.Param("id"))
	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewHandler("success", user))
}

func (controller *UsersController) Post(c Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	_, res := controller.Interactor.Create(id, name)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}

func (controller *UsersController) Put(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	_, res := controller.Interactor.Update(id, name)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}

func (controller *UsersController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, res := controller.Interactor.Delete(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewHandler(res.Error.Error(), nil))
		return
	}
	c.AbortWithStatus(res.StatusCode)
}
