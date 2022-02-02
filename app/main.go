package main

import (
	"github.com/44taka/golang-gin/infrastructure"
)

type User struct {
	ID   int
	Name string
}

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
