package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "password!"
	DBDbase = "cms"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error)
	}
	database = db
}
