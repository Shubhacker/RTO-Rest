package main

import (
	"log"
	"net/http"
	"os"
	apis "rto/APIs"
	"rto/database"

	"github.com/gorilla/mux"
)

func main() {

	database.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/login/", apis.Login).Methods("POST", "OPTIONS")
	// r.HandleFunc("/okay", PrintOkay).Methods("GET", "OPTIONS")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port : ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
