package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=qwerty12345 dbname=restaraunt_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to databse: ", err)
	}
	fmt.Println("connected to database successfully")
}
