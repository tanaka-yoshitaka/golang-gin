package usecase

import "github.com/44taka/golang-gin/domain"

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
