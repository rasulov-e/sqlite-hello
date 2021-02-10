package newsfeed

import "database/sql"

// Feed ...
type Feed struct {
	DB *sql.DB
}

func (feed *Feed) Add(item Item) {
	stmt, _ := feed.DB.Prepare(`
		INSERT INTO newsfeed (content) values (?)
	`)

	stmt.Exec(item.Content)

}

func NewFeed(db *sql.DB) *Feed {
	stmt, _ := db.Prepare(`
		CREATE TABLE "newsfeed" (
			"ID"	INTEGER NOT NULL,
			"content"	TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)

	stmt.Exec()

	return &Feed{
		DB: db,
	}
}
