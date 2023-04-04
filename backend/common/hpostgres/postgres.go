package hpostgres

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
func (r *PostgresConnection) Connect() error {
	db, err := sql.Open("postgres", r.ConnectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to postgres [ok]")

	r.Connected = true

	r.DB = db

	return nil
}

func (r *PostgresConnection) GetDB() (*sql.DB, error) {
	if r.DB == nil {
		err := r.Connect()
		if err != nil {
			return nil, err
		}
	}

	return r.DB, nil
}
