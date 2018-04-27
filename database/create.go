package database

var stmtsCreate = []string{
	"pragma auto_vacuum = full;",
	"create table vulnlist (id text not null primary key, title text, description text, uri text, creator text, impact text, solution text, date_public integer not null, date_publish integer not null, date_update integer not null);",
	"create table affected (id text not null, name text not null, product_name text not null, version_number text not null, primary key (id, name, product_name, version_number));",
	"create table cvss (id text not null, version text not null, base_vector text, base_score numeric, severity text, primary key (id, version));",
	"create table related (id text not null, type text, name text, vulinfo_id text not null, title text, url text, primary key (id, type, name, vulinfo_id));",
	"create table history (id text not null, history_no numeric not null, description text, date_time integer not null, primary key (id, history_no));",
	"create view vulnview as select list.id as id, list.title as title, list.description as description, list.uri as uri, list.impact as impact, list.solution as solution, cvss.base_score as cvss_score, cvss.severity as cvss_severity, list.date_public as date_public, list.date_publish as date_publish, list.date_update as date_update from vulnlist list left outer join cvss on list.id = cvss.id and cvss.version = '3.0';",
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

	for _, s := range stmtsCreate {
		err := func(st string) error {
			stmt, err := tx.Prepare(st)
			if err != nil {
				return err
			}
			defer stmt.Close()
			if _, err := stmt.Exec(); err != nil {
				return err
			}
			return nil
		}(s)
		if err != nil {
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
