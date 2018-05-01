package database

import "time"

//GetVulnview returns Vulnview instance
func (db *DB) GetVulnview(days int, score float64) ([]Vulnview, error) {
	ds := []Vulnview{}
	if db == nil {
		return ds, nil
	}
	logger := db.GetLogger()
	logger.Println("List JVN data:", db.GetDBFile())

	orderby := " order by date_update desc,id"
	if days > 0 {
		t := time.Now()
		start := time.Date(t.Year(), t.Month(), t.Day()-days, 0, 0, 0, 0, time.UTC)
		if _, err := db.GetDB().Select(&ds, "select * from vulnview where date_update >= ? and cvss_score >= ?"+orderby, start.Unix(), score); err != nil {
			return nil, err
		}
	} else {
		if _, err := db.GetDB().Select(&ds, "select * from vulnview where cvss_score >= ?"+orderby, score); err != nil {
			return nil, err
		}
	}
	return ds, nil
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
