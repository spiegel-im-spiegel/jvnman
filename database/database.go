package database

import (
	"database/sql"
	"io"
	"log"

	"github.com/go-gorp/gorp"
	myjvn "github.com/spiegel-im-spiegel/go-myjvn"
)

//DB is type of database
type DB struct {
	dbmap *gorp.DbMap
	api   *myjvn.APIs
}

//New returns DB instance
func New(dbf string, w io.Writer) (*DB, error) {
	db, err := sql.Open("sqlite3", dbf)
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	if w != nil {
		dbmap.TraceOn("[gorp]", log.New(w, "jvnman:", log.Lmicroseconds))
	} else {
		dbmap.TraceOff()
	}
	dbmap.AddTableWithName(Vulnlist{}, "vulnlist")
	dbmap.AddTableWithName(Affected{}, "affected")
	dbmap.AddTableWithName(CVSS{}, "cvss")
	dbmap.AddTableWithName(Related{}, "related")
	dbmap.AddTableWithName(History{}, "history")
	dbmap.AddTableWithName(Vulnview{}, "vulnview")

	return &DB{dbmap: dbmap, api: myjvn.New()}, nil
}

//GetDB returns sql.DB instance
func (db *DB) GetDB() *gorp.DbMap {
	if db == nil {
		return nil
	}
	return db.dbmap
}

//GetAPI returns myjvn.APIs instance
func (db *DB) GetAPI() *myjvn.APIs {
	if db == nil {
		return nil
	}
	return db.api
}
