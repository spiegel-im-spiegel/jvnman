package database

import (
	"database/sql"

	myjvn "github.com/spiegel-im-spiegel/go-myjvn"
)

//DB is type of database
type DB struct {
	db  *sql.DB
	api *myjvn.APIs
}

//New returns DB instance
func New(dbf string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbf)
	if err != nil {
		return nil, err
	}
	return &DB{db: db, api: myjvn.New()}, nil
}

//Close closes sql.DB
func (db *DB) Close() error {
	if db == nil {
		return nil
	}
	return db.db.Close()
}

//GetDB returns sql.DB instance
func (db *DB) GetDB() *sql.DB {
	if db == nil {
		return nil
	}
	return db.db
}
