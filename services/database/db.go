package database

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

// Connecting to db
func Init() {
	opts := &pg.Options{
		User:     "db_username",
		Password: "db_password",
		Addr:     "localhost:5432",
		Database: "db_dbname",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
}

func GetDB() *pg.DB {
	return db
}
