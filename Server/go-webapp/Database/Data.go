package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-webapp/Config"
	"log"
)

func Connection() (*sql.DB, error) {
	connStr := Config.Conn
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("An error has ocurred: %v", err)
	}

	return db, nil

}
