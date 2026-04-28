package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/clevextog/restaurant-api/internal/handler"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	connStr := "host=localhost port=5432 user=postgres password=qwerty12345 dbname=restaraunt_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to databse: ", err)
	}
	fmt.Println("connected to database successfully")

	http.HandleFunc("/", handler.Fallback)
	//http.Handle("/registration", PostOnlyMW(http.HandlerFunc(RegistrationHandler)))
	//http.Handle("/login", PostOnlyMW(http.HandlerFunc(LoginHandler)))
	//http.ListenAndServe(":8080", nil)
}
