package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sql/platform/newsfeed"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

/*
CREATE TABLE "newsfeed" (
	"ID"	INTEGER NOT NULL,
	"content"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
*/

var db *sql.DB
var dbPath string = "./newsfeed.db"

func initDB() {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := newsfeed.NewFeed(db)
	if err != nil {
		fmt.Printf("Error")
		return
	}

	item, _ := feed.Get()

}

func HandleAndServe() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.Handle("/", router)
}

func main() {
	initDB()

	HandleAndServe()
}
