package internals

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func TablesCreation(instanceOfDb *sql.DB) {
	var err error
	if instanceOfDb == nil {
		log.Println("Unable to reach the database")
		return
	}

	// Create the "Post" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Post (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			image TEXT DEFAULT NULL,
			title TEXT,
			content TEXT,
			user_id INTEGER,
			likes_count INTEGER DEFAULT 0,
			dislikes_count INTEGER DEFAULT 0,
			comments_count INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the "Session" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Session (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			token TEXT UNIQUE,
			expires TIMESTAMP,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the "User" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS User (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE,
			username TEXT UNIQUE,
			bio TEXT,
			avatar TEXT,
			password TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the "Category" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			category TEXT UNIQUE
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the "Category" table
	_, err = instanceOfDb.Exec(`
	INSERT OR IGNORE INTO Category (category) VALUES
    ('other'),
    ('music'),
    ('travel'),
    ('gaming'),
    ('sports'),
    ('programming'),
    ('politics'),
    ('entertainment'),
    ('movies'),
    ('books');

	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "Comments" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER,
			user_id INTEGER,
			comment TEXT,
			likes_count INTEGER DEFAULT 0,
			dislikes_count INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES Post(id),
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Create the junction table between post and category
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS PostCategories (
			post_id INTEGER,
			category_id INTEGER,
			FOREIGN KEY (post_id) REFERENCES Post (id),
			FOREIGN KEY (category_id) REFERENCES Category (id),
			PRIMARY KEY (post_id, category_id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// // Create the "Likes" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Commentlikes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			comment_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (comment_id) REFERENCES Comments(id),
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// // Create the "Likes" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Postlikes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES Post(id),
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Postdislikes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES Post(id),
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS Commentdislikes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			comment_id INTEGER,
			user_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (comment_id) REFERENCES Comments(id),
			FOREIGN KEY (user_id) REFERENCES User(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func DropAnTable(db *sql.DB, tableName string) {

	_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName))
	if err != nil {
		log.Fatal(err)
	}
}
