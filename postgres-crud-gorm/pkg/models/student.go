package models

import (
	"github.com/RiteshWanave/GO/postgres-crud-gorm/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name string
	PRN int	
}

func init () {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) Add () *Student {
	db.Create(&s)
	return s
}

func All () []Student {
	var students []Student
	db.Find(&students)
	return students	
}

func Get (id int) Student {
	var student Student
	db.Where("ID=?", id).Find(&student)
	return student
}

func Delete (id int) Student {
	var student Student
	db.Where("ID=?", id).Delete(student)
	return student
}

func (s *Student) Update (id int) Student {
	db.Where("ID=?", id).Save(&s)
	return *s
}