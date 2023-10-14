package main

import (
	"log"
	"net/http"
	"os"
	apis "rto/APIs"
	"rto/database"
)

func main() {

	database.ConnectDB()
	r := apis.Apis()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port : ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
