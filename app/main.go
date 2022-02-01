package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
type User struct {
	ID  int
	Name string
}

func main() {

	dsn := "root:root@tcp(mysql:3306)/my_testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error!!")
		return
	}
	var user User
	// var users []User
	fmt.Println(db.First(&user))
	// fmt.Println(db.Raw("select * from user limit 1").Scan(&user).Name)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"test": "test!!",
			"test2": "tesAAsdfasfasefaaaaaataaa",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
