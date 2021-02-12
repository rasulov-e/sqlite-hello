package newsfeed

import (
	"database/sql"
	"log"
)

// Feed ...
type Feed struct {
	DB *sql.DB
}

func (feed *Feed) Get() ([]Item, error) {
	items := []Item{}

	rows, err := feed.DB.Query(`
		SELECT * FROM newsfeed
	`)

	if err != nil {
		log.Printf("newsfeed: NewFeed: %s", err.Error())
		return nil, err
	}

	var id int
	var content string

	for rows.Next() {
		err = rows.Scan(&id, &content)
		if err != nil {
			log.Printf("newsfeed: Get: %s", err.Error())
		}
		item := Item{
			ID:      id,
			Content: content,
		}
		items = append(items, item)
	}
	return items, nil
}

func (feed *Feed) Add(item Item) (int64, error) {
	stmt, _ := feed.DB.Prepare(`
		INSERT INTO newsfeed (content) values (?)
	`)

	res, err := stmt.Exec(item.Content)
	id, _ := res.LastInsertId()

	if err != nil || id == 0 {
		log.Printf("newsfeed: Add: error")
		return id, err
	}
	return id, nil
}

func NewFeed(db *sql.DB) (*Feed, error) {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "newsfeed" (
			"ID"	INTEGER NOT NULL,
			"content"	TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
		);
	`)
	if err != nil {
		log.Printf("newsfeed: NewFeed: %s", err.Error())
		return nil, err
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Printf("newsfeed: NewFeed: %s", err.Error())
		return nil, err
	}

	return &Feed{
		DB: db,
	}, nil
}
