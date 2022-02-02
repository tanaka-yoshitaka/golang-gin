package infrastructure

import (
	"github.com/44taka/golang-gin/infrastructure/middleware"
	"github.com/44taka/golang-gin/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB, middleware ...gin.HandlerFunc) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setMiddleware()
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	// 疎通確認
	r.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// ユーザー周り
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
	r.Gin.POST("/users", func(c *gin.Context) { usersController.Post(c) })
	r.Gin.PUT("/users/:id", func(c *gin.Context) { usersController.Put(c) })
	r.Gin.DELETE("/users/:id", func(c *gin.Context) { usersController.Delete(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}

func (r *Routing) setMiddleware() {
	// TODO::ミドルウェアのセットをもっと効率よくしたい
	r.Gin.Use(middleware.MyCustomLogger())
	r.Gin.Use(middleware.RequestID())
}
