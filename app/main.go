package main

import (
	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/infrastructure/persistence"
	"github.com/44taka/golang-gin/interfaces/handler"
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	// コンフィグ読み込み
	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)

	userPersistence := persistence.NewUserPersistence(db.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) { userHandler.FindAll(ctx) })
	r.GET("/users/:id", func(ctx *gin.Context) { userHandler.FindById(ctx) })
	r.POST("/users", func(ctx *gin.Context) { userHandler.Create(ctx) })
	r.PUT("/users/:id", func(ctx *gin.Context) { userHandler.Update(ctx) })
	r.DELETE("/users/:id", func(ctx *gin.Context) { userHandler.Delete(ctx) })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
