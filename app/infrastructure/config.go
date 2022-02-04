package infrastructure

import (
	"golang.org/x/text/language"
)

type Config struct {
	DB struct {
		Local struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
	Language []language.Tag
	Jwt      struct {
		SecretKey       string
		Issuer          string
		ExpirationHours int
	}
}

func NewConfig() *Config {
	c := new(Config)

	// DB設定
	c.DB.Local.Host = "mysql"
	c.DB.Local.Username = "root"
	c.DB.Local.Password = "root"
	c.DB.Local.DBName = "my_testdb"

	// ポート番号
	c.Routing.Port = ":8080"

	// 対応言語
	c.Language = []language.Tag{
		// 対応する言語はここに追記していく
		language.Japanese,
		language.English,
	}

	// JWT関連
	c.Jwt.SecretKey = "verysecretkey"
	c.Jwt.Issuer = "golang-gin"
	c.Jwt.ExpirationHours = 1

	return c
}
