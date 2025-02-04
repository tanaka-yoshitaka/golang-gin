package usecase

import (
	"net/http"

	"github.com/44taka/golang-gin/domain"
	"github.com/44taka/golang-gin/utils"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{}, NewResultStatus(http.StatusNotFound, err)
	}
	user = foundUser.BuildForGet()

	return user, NewResultStatus(http.StatusOK, nil)
}

func (interactor *UserInteractor) Create(id int, name string) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Create(db, id, name)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusCreated, nil)
}

func (interactor *UserInteractor) Update(id int, name string) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Update(db, id, name)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusNoContent, nil)
}

func (interactor *UserInteractor) Delete(id int) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Delete(db, id)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusNoContent, nil)
}

func (interactor *UserInteractor) Login(id int, password string) (user domain.Users, token string, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Login(db, id, password)
	if err != nil {
		return domain.Users{}, "", NewResultStatus(http.StatusBadRequest, err)
	}

	// TODO::シークレットキーなどは設定ファイルから読み込ませる
	jwtWrppaer := utils.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	token, err = jwtWrppaer.GenerateToken(user.ID)

	return user, token, NewResultStatus(http.StatusOK, nil)
}
