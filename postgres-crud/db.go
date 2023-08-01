package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init () {
	var err error
	db, err = sql.Open("postgres", "user=ritesh password=vroot dbname=db1 sslmode=disable")
	CheckErr(err)
}

func Add (Name string, Roll_Number int) {
	insert := "INSERT INTO students (Name, Roll_Number) VALUES ($1, $2)"
	_, err := db.Exec(insert, Name, Roll_Number)
	CheckErr(err)
	fmt.Printf("Student with name %s and roll number %d added\n", Name, Roll_Number)
}

func Delete (Roll_Number int) {
	delete := "DELETE FROM students WHERE Roll_Number = $1"
	_, err := db.Exec(delete, Roll_Number)
	CheckErr(err)
	fmt.Printf("Student with roll number %d deleted\n", Roll_Number)
}

func Update (Roll_Number int) {
	var Name string
	fmt.Print("Enter the new name: ")
	fmt.Scan(&Name)

	update := "UPDATE students SET Name = $1 WHERE Roll_Number = $2"
	_, err := db.Exec(update, Name, Roll_Number)
	CheckErr(err)
	fmt.Printf("Student with roll number %d updated\n", Roll_Number)
}

func Get (Roll_Number int) {
	get := "SELECT Name FROM students WHERE Roll_Number = $1"
	rows, err := db.Query(get, Roll_Number)
	CheckErr(err)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		CheckErr(err)
		println(name)
	}
}

func GetAll () {
	getAll := "SELECT * FROM students"
	rows, err := db.Query(getAll)
	CheckErr(err)
	for rows.Next() {
		var name string
		var roll_number int
		err = rows.Scan(&name, &roll_number)
		CheckErr(err)
		println(name, roll_number)
	}
}