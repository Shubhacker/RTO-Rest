package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

const (
	HOST     = "containers-us-west-176.railway.app"
	PORT     = 5913
	USER     = "postgres"
	PASSWORD = "S2SrkFGkTXyH2TxPMD57"
	DBNAME   = "railway"
)

func ConnectDB() *sql.DB {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully !...")
	return DB
	// defer DB.Close()
}
