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
	// コンフィグ読み込み
	c := infrastructure.NewConfig()
	// DB接続情報読み込み
	db := infrastructure.NewDB(c)
	// ルーティング設定読み込み
	r := infrastructure.NewRouting(db, c)

	r.Run()
}
