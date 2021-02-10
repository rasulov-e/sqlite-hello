package main

import (
	"database/sql"
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
	db, _ := sql.Open("sqlite3", "./newsfeed.db")
	feed := newsfeed.NewFeed(db)

	feed.Add(newsfeed.Item{
		Content: "Hello! SQL",
	})
}
