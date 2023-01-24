package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// getConnection get a connection to the DB
func GetConnection() *sql.DB {
	var err error
	dsn := os.Getenv("DB_URL")
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
