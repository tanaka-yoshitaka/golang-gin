package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	return newDB(&DB{
		Host:     "mysql",
		Username: "root",
		Password: "root",
		DBName:   "my_testdb",
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open(mysql.Open(d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
