package main

import (
	"database/sql"
	"fmt"
	"log"
	"sql/platform/newsfeed"

	_ "github.com/mattn/go-sqlite3"
)

/*
CREATE TABLE "newsfeed" (
	"ID"	INTEGER NOT NULL,
	"content"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
*/

func main() {
	db, err := sql.Open("sqlite3", "./newsfeed.db")
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	feed, err := newsfeed.NewFeed(db)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	item, _ := feed.Get()

	fmt.Println(item)
}
