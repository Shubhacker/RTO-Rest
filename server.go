package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login/", Login).Methods("POST", "OPTIONS")
	// r.HandleFunc("/okay", PrintOkay).Methods("GET", "OPTIONS")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port : ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

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
