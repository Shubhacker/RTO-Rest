package main

import (
	"log"
	"net/http"
	"os"
	apis "rto/APIs"
	"rto/database"
)

func main() {
	log.Println("Starting Server")
	database.ConnectDB()
	r := apis.Apis()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	l1 := "Server running on port : " + port
	log.Println(l1)
	go database.Logs("serverStarting", l1)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
