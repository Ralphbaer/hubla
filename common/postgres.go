package common

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// PostgresConnection is a hub which deal with mongodb connections.
type PostgresConnection struct {
	ConnectionString string
	DB               *sql.DB
	Connected        bool
}

// Connect keeps a singleton connection with mongodb.
func (r *PostgresConnection) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to postgres [ok]")

	r.Connected = true

	return db, nil
}
