package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect () {
	d, err := gorm.Open("mysql", "ritesh:vroot@tcp(127.0.0.1:3306)/goBookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db = d
	fmt.Println("Connected to database 📑")
}

func GetDB () *gorm.DB {
	return db
}