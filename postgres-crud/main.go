package main

import (
	"flag"
)

func main () {
	newName := flag.String("name", "", "Name of the student")
	newRollNumber := flag.Int("roll", 0, "Roll number of the student")
	delete := flag.Int("delete", 0, "Roll number of the student to be deleted")
	update := flag.Int("update", 0, "Roll number of the student to be updated")
	getAll := flag.Bool("all", false, "Get all the students")
	get := flag.Int("get", 0, "Roll number of the student to be fetched")
	flag.Parse()

	switch {
	case *newName != "" && *newRollNumber != 0:
		Add(*newName, *newRollNumber)

	case *delete != 0:
		Delete(*delete)

	case *update != 0:
		Update(*update)

	case *getAll:
		GetAll()

	case *get != 0:
		Get(*get)
		
	default:
		flag.PrintDefaults()
	}
}