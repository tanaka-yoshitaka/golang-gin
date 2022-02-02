package main

import (
	"fmt"

	"github.com/44taka/golang-gin/infrastructure"
)

type User struct {
	ID   int
	Name string
}

func main() {
	// asdfasnjhkjhk
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	var user User
	fmt.Println(db.Connection.First(&user))
	r.Run()
}
