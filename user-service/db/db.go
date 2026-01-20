// user-service/db/db.go
package db

import (
	"database/sql"
	"log"
	"time"
)

func NewDB() *sql.DB {
	var err error

	// Retry logic
	for i := 0; i < 10; i++ { // try 10 times
		db, err := sql.Open("sqlite", "users.db")
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("✅ Connected to PostgreSQL")
				return db
			}
		}
		log.Printf("⚠️ Sqlite not ready yet, retrying in 3s... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}
	log.Fatal("❌ Could not connect to Sqlite:", err)
	return nil

}
