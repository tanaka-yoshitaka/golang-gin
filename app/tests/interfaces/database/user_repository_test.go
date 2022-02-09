package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/44taka/golang-gin/domain"
	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func TestGetAll(t *testing.T) {
	// c := infrastructure.NewConfig()
	// db := infrastructure.NewDB(c)
	// usersController := controllers.NewUsersController(db)

	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	c := infrastructure.NewConfig()
	db := infrastructure.NewDB(c)
	usersController := controllers.NewUsersController(db)
	usersController.GetAll(ctx)
	fmt.Println("--------------")
	var usersResult domain.UsersResult
	_ = json.Unmarshal(response.Body.Bytes(), &usersResult)
	fmt.Println(response.Body.Bytes())
	fmt.Println(usersResult.Status)
	fmt.Println("--------------")
}
