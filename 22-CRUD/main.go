package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcosrosse/golang-course/22-CRUD/server"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Listening in the port 8080")

	router.HandleFunc("/users", server.CreateUsers).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
