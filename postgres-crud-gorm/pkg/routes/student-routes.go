package routes

import (
	"github.com/RiteshWanave/GO/postgres-crud-gorm/pkg/controllers"
	"github.com/gorilla/mux"
)


var StudentRouter = func (router *mux.Router) {

	router.HandleFunc("/student", controllers.GetAll).Methods("GET")
	router.HandleFunc("/student", controllers.AddStudent).Methods("POST")
	router.HandleFunc("/student/{id}", controllers.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/student/{id}", controllers.GetById).Methods("GET")
	router.HandleFunc("/student/{id}", controllers.UpdateStudent).Methods("PUT")
}