package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := "root:password@/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
