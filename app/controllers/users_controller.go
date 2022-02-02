package controllers

import (
	"strconv"

	"github.com/44taka/golang-gin/interfaces/database"
	"github.com/44taka/golang-gin/usecase"
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
	id, _ := strconv.Atoi(c.Param("id"))
	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.Status, NewHandler(res.Error.Error(), nil))
	}
	c.JSON(res.Status, NewHandler("success", user))
}
