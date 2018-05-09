package database

import (
	"time"

	"github.com/Masterminds/squirrel"
)

//GetVulnview returns Vulnview instance
func (db *DB) GetVulnview(days int, score float64) ([]Vulnview, error) {
	ds := []Vulnview{}
	if db == nil {
		return ds, nil
	}
	builder := squirrel.Select("*").From("vulnview").OrderBy("date_update desc", "id")
	and := squirrel.And{}
	and = append(and, squirrel.GtOrEq{"cvss_score": score})
	if days > 0 {
		t := time.Now()
		start := time.Date(t.Year(), t.Month(), t.Day()-days, 0, 0, 0, 0, time.UTC)
		and = append(and, squirrel.GtOrEq{"date_update": start.Unix()})
	}
	if psql, args, err := builder.Where(and).ToSql(); err != nil {
		return ds, err
	} else if _, err = db.GetDB().Select(&ds, psql, args); err != nil {
		return ds, err
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
