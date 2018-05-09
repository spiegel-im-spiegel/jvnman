package database

import (
	"fmt"
	"time"

	"gopkg.in/Masterminds/squirrel.v1"
)

//GetVulnview returns Vulnview instance
func (db *DB) GetVulnview(days int, score float64, product string) ([]Vulnview, error) {
	ds := []Vulnview{}
	if db == nil {
		return ds, nil
	}
	bldr := squirrel.Select(
		"vulnview.id as id",
		"vulnview.title as title",
		"vulnview.description as description",
		"vulnview.uri as uri",
		"vulnview.impact as impact",
		"vulnview.solution as solution",
		"vulnview.cvss_score as cvss_score",
		"vulnview.cvss_severity as cvss_severity",
		"vulnview.date_public as date_public",
		"vulnview.date_publish as date_publish",
		"vulnview.date_update as date_update ",
	).From("vulnview").OrderBy("date_update desc", "id")

	//make "where" condition
	if days > 0 {
		t := time.Now()
		bldr = bldr.Where(squirrel.GtOrEq{"date_update": time.Date(t.Year(), t.Month(), t.Day()-days, 0, 0, 0, 0, time.UTC).Unix()})
	}
	if score > 0 {
		bldr = bldr.Where(squirrel.GtOrEq{"cvss_score": score})
	}
	if len(product) > 0 {
		sq, args, err := squirrel.Select("id").Distinct().From("affected").Where(squirrel.Expr("product_name like ?", "%"+product+"%")).ToSql()
		if err != nil {
			return ds, err
		}
		bldr = bldr.Join(fmt.Sprintf("(%s) affect on affect.id = vulnview.id", sq), args...)
	}

	//query
	if psql, args, err := bldr.ToSql(); err != nil {
		return ds, err
	} else if _, err = db.GetDB().Select(&ds, psql, args...); err != nil {
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
