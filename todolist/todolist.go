package main

import (
	"database/sql"
	"fmt"


	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/madhvip/golang_projects/todolist/router"
)

const (
	DB_USER = "madhvip"
	DB_PASS = "madhvip"
	DB_NAME = "friends_repo"
)

var db *sql.DB
var err error

func init() {
	log.Info("Connecting to todolist database")
	db, err = connectToDatabase()
	chkerr(err)
}

func main() {
	defer db.Close()
	log.Info("Starting todolist server")
	
	http.ListenAndServe(":9000", nil)

}

func connectToDatabase() (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("host=localhost port= 5432 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	return
}

func chkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}
