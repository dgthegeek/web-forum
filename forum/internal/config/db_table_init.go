package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var predefinedCategories = []string{
	"technology",
	"sport",
	"finance",
}

func CreateTables(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			email TEXT UNIQUE,
			password TEXT
		);
	`)
	if err != nil {
		log.Fatal("user table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255),
			content TEXT,
			author VARCHAR(255),
			createdAt DATETIME,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0
		);
	`)
	if err != nil {
		log.Fatal("posts table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255)
		);
	`)
	if err != nil {
		log.Fatal("categories table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS post_categories (
			postID INTEGER,
			categoryID INTEGER,
			FOREIGN KEY (postID) REFERENCES posts(id),
			FOREIGN KEY (categoryID) REFERENCES categories(id),
			PRIMARY KEY (postID, categoryID)
		);
	`)
	if err != nil {
		log.Fatal("post_categories table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS postLikes (
			id INTEGER PRIMARY KEY,
			postID INTEGER,
			userID INTEGER,
			FOREIGN KEY (postID) REFERENCES posts(id),
			FOREIGN KEY (userID) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal("likes table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			idComment INTEGER PRIMARY KEY AUTOINCREMENT,
			commentText TEXT,
			username TEXT,
			nbrlikes INTEGER DEFAULT 0,
			nbrdislikes INTEGER DEFAULT 0,
			idPost INTEGER,
			FOREIGN KEY (idPost) REFERENCES posts (id)
		);
	`)
	if err != nil {
		log.Fatal("Comments tables creation error:",err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS postDislikes (
			id INTEGER PRIMARY KEY,
			postID INTEGER,
			userID INTEGER,
			FOREIGN KEY (postID) REFERENCES posts(id),
			FOREIGN KEY (userID) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal("dislikes table creation error:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS commentDislikes (
			id INTEGER PRIMARY KEY,
			commentID INTEGER,
			userID INTEGER,
			FOREIGN KEY (commentID) REFERENCES comments(idComment),
			FOREIGN KEY (userID) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal("dislikes of comments table creation error:", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS commentLikes (
			id INTEGER PRIMARY KEY,
			commentID INTEGER,
			userID INTEGER,
			FOREIGN KEY (commentID) REFERENCES comments(idComment),
			FOREIGN KEY (userID) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal("likes of comments table creation error:", err)
	}

	// Insert predefined category data
	insertCategoryStmt, err := db.Prepare("INSERT INTO categories (name) VALUES (?)")
	if err != nil {
		log.Fatal("category data insertion error:", err)
	}
	defer insertCategoryStmt.Close()

	for _, categoryName := range predefinedCategories {
		_, err = insertCategoryStmt.Exec(categoryName)
		if err != nil {
			log.Fatal("category data insertion error:", err)
		}
	}

	fmt.Println("db created")

}
