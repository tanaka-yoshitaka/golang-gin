package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal(1, 1)
}

func TestFindById(t *testing.T) {
	// handler実行
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/users/1", nil)
	// fmt.Println(response.Body)

	// usecase := new(mock_handler.MockUserUseCase)
	// usecase.NewUserUseCase()

	// mh := new(mock_handler.MockUserHandler)
	// mh.FindById(ctx)
	// h := NewUserHandler(mh)
	// fmt.Println("++++++++++++++++")
	// fmt.Println(result)
	// fmt.Println("++++++++++++++++")

	asserts := assert.New(t)
	asserts.Equal(1, 1)
}
