package internals

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Driver   string
	Name     string
	PORT     string
	Username string
	Password string
}

/*
Instanciate a new Database
in case of SQlite connection only the driver and
the name are needed the name corresponds to the path of the database file
*/
func (db *Config) Init() (*sql.DB, error) {
	// Check if the database file exists in root directory

	// Etablish a new database connection
	databaseConnection, err := sql.Open(db.Driver, db.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return databaseConnection, nil
}
