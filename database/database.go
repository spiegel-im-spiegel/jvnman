package database

import (
	"database/sql"

	"gopkg.in/gorp.v2"
	myjvn "github.com/spiegel-im-spiegel/go-myjvn"
	"github.com/spiegel-im-spiegel/logf"
)

//DB is type of database
type DB struct {
	dbf    string
	dbmap  *gorp.DbMap
	api    *myjvn.APIs
	logger *logf.Logger
}

//New returns DB instance
func New(dbf string, logger *logf.Logger) (*DB, error) {
	db, err := sql.Open("sqlite3", dbf)
	if err != nil {
		logger.Fatalln(err)
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	if logger.MinLevel() <= logf.TRACE {
		dbmap.TraceOn("[TRACE]", logger.GetLogger())
	} else {
		dbmap.TraceOff()
	}
	dbmap.AddTableWithName(Vulnlist{}, "vulnlist")
	dbmap.AddTableWithName(Affected{}, "affected")
	dbmap.AddTableWithName(CVSS{}, "cvss")
	dbmap.AddTableWithName(Related{}, "related")
	dbmap.AddTableWithName(History{}, "history")
	dbmap.AddTableWithName(Vulnview{}, "vulnview")

	return &DB{dbf: dbf, dbmap: dbmap, api: myjvn.New(), logger: logger}, nil
}

//GetDBFile returns SQLite file path
func (db *DB) GetDBFile() string {
	if db == nil {
		return ""
	}
	return db.dbf
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

//GetLogger returns logf.Logger instance
func (db *DB) GetLogger() *logf.Logger {
	if db == nil {
		return nil
	}
	return db.logger
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
