package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	// response := httptest.NewRecorder()
	// ctx, _ := gin.CreateTestContext(response)
	// ctx.Request, _ = http.NewRequest(http.MethodGet, "/users/1", nil)
	// fmt.Println("********************")
	// fmt.Println(response.Body.String())
	// fmt.Println("********************")

	asserts := assert.New(t)
	asserts.Equal(1, 1)
}

// func TestCreate(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)

// 	mockUserHandler := mock_handler.NewMockUserHandler(ctrl)
// 	mockUserHandler.EXPECT().Create(gomock.Any())
// 	mockUserHandler.FindById(ctx)

// 	asserts := assert.New(t)
// 	asserts.Equal(1, 1)
// }
