package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login/", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}

type response struct {
	Message   string
	Username  string
	FirstName string
	LastName  string
	UserId    string
	JWT       string
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Login API")

	var responseData response

	json.NewEncoder(w).Encode(responseData)

}
