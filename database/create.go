package database

var stmts = []string{
	"create table vulnlist (id text not null primary key, title text, description text, uri text, creator text, date_public text, date_publish text, date_update text);",
}

//Initialize returns result for initializing
func (db *DB) Initialize() error {
	if db == nil {
		return nil
	}
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	for _, s := range stmts {
		stmt, err := tx.Prepare(s)
		if err != nil {
			tx.Rollback()
			return err
		}
		if _, err := stmt.Exec(); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
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
