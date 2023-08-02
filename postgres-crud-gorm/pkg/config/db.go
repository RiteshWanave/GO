package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect () {
	d, err := gorm.Open("postgres", "user=ritesh password=vroot dbname=db2 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d
	println("Connected to database 📑")
}

func GetDB() *gorm.DB {
	return db
}