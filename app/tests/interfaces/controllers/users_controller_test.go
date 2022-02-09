package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/44taka/golang-gin/domain"
	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// 期待値の生成
	expected := map[string]interface{}{
		"status":  200,
		"message": "OK",
		"result": []map[string]interface{}{
			{
				"id":   1,
				"name": "ichiro",
			},
			{
				"id":   2,
				"name": "jiro",
			},
			{
				"id":   3,
				"name": "saburo",
			},
			{
				"id":   4,
				"name": "shiro",
			},
			{
				"id":   5,
				"name": "goro",
			},
			{
				"id":   6,
				"name": "rokuro",
			},
			{
				"id":   7,
				"name": "shichiro",
			},
			{
				"id":   8,
				"name": "hachiro",
			},
			{
				"id":   9,
				"name": "kuro",
			},
			{
				"id":   10,
				"name": "juro",
			},
		},
	}
	expectedByte, _ := json.Marshal(expected)
	var expectedUsersResult domain.UsersResult
	json.Unmarshal(expectedByte, &expectedUsersResult)

	// コントローラの実行
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/auth/users", nil)

	c := infrastructure.NewConfig()
	db := infrastructure.NewDB(c)
	usersController := controllers.NewUsersController(db)
	usersController.GetAll(ctx)

	var actualUsersResult domain.UsersResult
	json.Unmarshal(response.Body.Bytes(), &actualUsersResult)

	// 比較
	asserts := assert.New(t)
	asserts.Equal(expectedUsersResult, actualUsersResult)
}

// func TestGet(t *testing.T) {
// 	// 期待値の生成
// 	expected := map[string]interface{}{
// 		"status":  200,
// 		"message": "OK",
// 		"result": []map[string]interface{}{
// 			{
// 				"id":   1,
// 				"name": "ichiro",
// 			},
// 		},
// 	}
// 	expectedByte, _ := json.Marshal(expected)
// 	var expectedUsersResult domain.UsersResult
// 	json.Unmarshal(expectedByte, &expectedUsersResult)

// 	// コントローラの実行
// 	response := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(response)
// 	// TODO::パスパラメータがコントローラにうまく渡らない。。
// 	ctx.Request, _ = http.NewRequest(http.MethodGet, "/auth/users", nil)
// 	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	c := infrastructure.NewConfig()
// 	db := infrastructure.NewDB(c)
// 	usersController := controllers.NewUsersController(db)
// 	usersController.Get(ctx)

// 	var actualUsersResult domain.UsersResult
// 	json.Unmarshal(response.Body.Bytes(), &actualUsersResult)

// 	// 比較
// 	asserts := assert.New(t)
// 	asserts.Equal(expectedUsersResult, actualUsersResult)
// }

func TestPost(t *testing.T) {
	// コントローラの実行
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest(
		http.MethodPost,
		"/auth/users",
		strings.NewReader("name=testarou&password=password"),
	)
	ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c := infrastructure.NewConfig()
	db := infrastructure.NewDB(c)
	usersController := controllers.NewUsersController(db)
	usersController.Post(ctx)

	// 比較
	asserts := assert.New(t)
	asserts.Equal(http.StatusCreated, response.Result().StatusCode)
}
