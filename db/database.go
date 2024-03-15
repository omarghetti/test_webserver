package db

import (
	"log"

	"github.com/go-pg/pg"
)

func ConnectToDatabase() *pg.DB {
	// Connect to the database
	log.Printf("Connecting to the database at address: %s", "localhost:5432")
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "Bellali95",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
	defer db.Close()
	return db
}
