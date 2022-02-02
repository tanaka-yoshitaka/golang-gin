package infrastructure

import (
	"github.com/44taka/golang-gin/controllers"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	r.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"test":    "test!!",
			"test2":   "tesAAsasdfaefdfasfasefaaaaaataaaaaafkmawoefj;aowiejf;oaiwejaaa",
		})
	})
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
