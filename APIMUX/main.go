package main

import (
	"database/sql"
	"log"
	"net/http"

	"Semester4/APIMUX/Controller"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	r := mux.NewRouter()

	db = Controller.Connect()

	r.HandleFunc("/create", Controller.CreateHandler).Methods("POST")
	r.HandleFunc("/retrieve/{id}", Controller.RetrieveHandler).Methods("GET")
	r.HandleFunc("/update/{id}", Controller.UpdateHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", Controller.DeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
