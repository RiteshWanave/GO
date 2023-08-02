package main

import (
	"log"
	"net/http"

	"github.com/RiteshWanave/GO/postgres-crud-gorm/pkg/routes"
	"github.com/gorilla/mux"
)

func main () {
	router := mux.NewRouter()
	routes.StudentRouter(router)
	http.Handle("/", router)

	println("Server is running on port 6000 ✅")
	log.Fatal(http.ListenAndServe("localhost:6000", router))
}