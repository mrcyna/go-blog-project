package db

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Post struct {
	ID        string
	Title     string
	Slug      string
	Body      string
	Author    string
	Published bool
}

func init() {

	var err error

	// Database Connection
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/blog")
	if err != nil {
		panic(err)
	}

	// Database Connection Settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Create Posts Table
	_, err = db.Query(`
			CREATE TABLE IF NOT EXISTS posts (
				id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
				title VARCHAR(128) NOT NULL,
				slug VARCHAR(128) NOT NULL,
				body TEXT,
				AUTHOR VARCHAR(64),
				published TINYINT
			);
		`)
	if err != nil {
		panic(err)
	}
}

func CreatePost(title, body, author string, published bool) bool {

	// Title
	title = strings.ReplaceAll(title, "\n", "")

	// Slug
	slug := strings.ReplaceAll(title, " ", "-")
	slug = strings.ToLower(slug)

	// Body
	body = strings.ReplaceAll(body, "\n", "")

	// Author
	author = strings.ReplaceAll(author, "\n", "")

	// SQL Statement
	stmt, err := db.Prepare("INSERT INTO posts(title, slug, body, author, published) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	// Query Execution
	res, err := stmt.Exec(title, slug, body, author, published)
	if err != nil {
		panic(err.Error())
	}

	_, err = res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return true
}
