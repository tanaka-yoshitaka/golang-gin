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

// 全てユーザー情報取得
func (interactor *UserInteractor) GetAll() (users []domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	users, err := interactor.User.FindAll(db)
	if err != nil {
		return users, NewResultStatus(http.StatusNotFound, err)
	}

	return users, NewResultStatus(http.StatusOK, nil)
}

// ユーザー情報取得
func (interactor *UserInteractor) Get(id int) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.FindByID(db, id)
	if err != nil {
		return user, NewResultStatus(http.StatusNotFound, err)
	}

	return user, NewResultStatus(http.StatusOK, nil)
}

// ユーザー新規作成
func (interactor *UserInteractor) Create(id int, name string) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Create(db, id, name)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusCreated, nil)
}

// ユーザー情報更新
func (interactor *UserInteractor) Update(id int, name string) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Update(db, id, name)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusNoContent, nil)
}

// ユーザー情報削除
func (interactor *UserInteractor) Delete(id int) (user domain.Users, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Delete(db, id)
	if err != nil {
		return domain.Users{}, NewResultStatus(http.StatusBadRequest, err)
	}
	return user, NewResultStatus(http.StatusNoContent, nil)
}

// ログインユーザー情報取得
func (interactor *UserInteractor) Login(id int, password string) (user domain.Users, token string, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.Login(db, id, password)
	if err != nil {
		return domain.Users{}, "", NewResultStatus(http.StatusBadRequest, err)
	}

	// TODO::シークレットキーなどは設定ファイルから読み込ませる
	// アクセストークン発行
	jwtWrppaer := utils.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 1,
	}
	token, err = jwtWrppaer.GenerateToken(user.ID)
	if err != nil {
		return domain.Users{}, "", NewResultStatus(http.StatusBadRequest, err)
	}

	return user, token, NewResultStatus(http.StatusOK, nil)
}
