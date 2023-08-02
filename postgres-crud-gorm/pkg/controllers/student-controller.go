package controllers

import (
	"net/http"
	"strconv"

	"github.com/RiteshWanave/GO/postgres-crud-gorm/pkg/models"
	"github.com/RiteshWanave/GO/postgres-crud-gorm/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAll (w http.ResponseWriter, r *http.Request) {
	student := models.All()
	utils.ResponseWriter(w, student)
}

func GetById (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		println("error while parsing")
	}
	student := models.Get(int(id))
	utils.ResponseWriter(w, student)
}

func AddStudent (w http.ResponseWriter, r *http.Request) {
	student := &models.Student{}
	utils.ParseBody(r, student)
	newstudent := student.Add()
	utils.ResponseWriter(w, newstudent)
}

func DeleteStudent (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		println("error while parsing")
	}
	student := models.Delete(int(id))
	utils.ResponseWriter(w, student)
}

func UpdateStudent (w http.ResponseWriter, r *http.Request) {
	updatedstudent := &models.Student{}
	utils.ParseBody(r, updatedstudent)
	vars := mux.Vars(r)
	ID := vars["id"]
	id, err := strconv.ParseInt(ID, 0, 0)
	if err != nil {
		println("error while parsing")
	}
	student := updatedstudent.Update(int(id))
	utils.ResponseWriter(w, student)
}