package infrastructure

import (
	"net/http"

	"github.com/44taka/golang-gin/infrastructure/middleware"
	"github.com/44taka/golang-gin/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB, c *Config) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting(c)
	return r
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func (r *Routing) setRouting(c *Config) {
	// 共通ミドルウェア設定
	r.Gin.Use(middleware.MyCustomLogger())
	r.Gin.Use(middleware.SetLaguageMessage(c.Language))
	r.Gin.Use(middleware.RequestID())

	// 疎通確認
	r.Gin.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	// ユーザー周り
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.POST("/login", func(c *gin.Context) { usersController.Login(c) })

	// 認証グループ
	authGroup := r.Gin.Group("/auth", middleware.AuthMiddleware())
	authGroup.POST("/users", func(c *gin.Context) { usersController.Post(c) })
	authGroup.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
	authGroup.PUT("/users/:id", func(c *gin.Context) { usersController.Put(c) })
	authGroup.DELETE("/users/:id", func(c *gin.Context) { usersController.Delete(c) })
	authGroup.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "auth_pong"}) })

	// 404
	r.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "page not found.",
		})
	})
	// 405
	r.Gin.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusMethodNotAllowed,
			"message": "method not allowed.",
		})
	})

}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
