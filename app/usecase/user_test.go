package usecase

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/44taka/golang-gin/domain/model"
	"github.com/44taka/golang-gin/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestUseCaseUserFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := model.User{
		ID:       2,
		Name:     "jiro",
		Password: "password",
	}
	var err error

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	mockUserRepository := mocks.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindById(ctx, 2).Return(expected, err)

	userUseCase := NewUserUseCase(mockUserRepository)
	result, err := userUseCase.FindById(ctx, 2)
	if err != nil {
		t.Error("Error!!!!!")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindById() is not same as expected")
	}

	// fmt.Println("************")
	// fmt.Println(result)
	// fmt.Println(expected)
	// fmt.Println(err.Error())
	// fmt.Println("************")

	// asserts := assert.New(t)
	// if err != nil {
	// 	asserts.Error(err)
	// } else {
	// 	asserts.Equal(1, 1)
	// }
}
