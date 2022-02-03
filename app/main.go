package main

import (
	"github.com/44taka/golang-gin/infrastructure"
)

func init() {
	// 言語ファイル読み込み
	c := infrastructure.NewConfig()
	infrastructure.NewInitLanguage(c)
}

func main() {
	c := infrastructure.NewConfig()
	db := infrastructure.NewDB(c)
	r := infrastructure.NewRouting(db, c)
	r.Run()
}
