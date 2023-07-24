package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/RiteshWanave/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main () {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Server is running at port 5000 ✅")
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}