package database

import (
	"fmt"
	"time"

	"gopkg.in/Masterminds/squirrel.v1"
)

//GetVulnview returns Vulnview instance
func (db *DB) GetVulnview(id string) *Vulnview {
	ds := Vulnview{}
	if psql, args, err := selectVulnview.Where(squirrel.Eq{"id": id}).ToSql(); err != nil {
		db.GetLogger().Println(err)
		return nil
	} else if err := db.GetDB().SelectOne(&ds, psql, args...); err != nil {
		db.GetLogger().Println(err)
		return nil
	}
	return &ds
}

//GetVulnviewList returns []Vulnview instance
func (db *DB) GetVulnviewList(days int, score float64, product, cve string) ([]Vulnview, error) {
	ds := []Vulnview{}
	if db == nil {
		return ds, nil
	}
	bldr := selectVulnview

	//make "where" condition
	if days > 0 {
		t := time.Now()
		bldr = bldr.Where(squirrel.GtOrEq{"date_update": time.Date(t.Year(), t.Month(), t.Day()-days, 0, 0, 0, 0, time.UTC).Unix()})
	}
	if score > 0 {
		bldr = bldr.Where(squirrel.GtOrEq{"cvss_score": score})
	}
	if len(product) > 0 {
		likeProduct := "%" + product + "%"
		cond := squirrel.Or{squirrel.Expr("name like ?", likeProduct), squirrel.Expr("product_name like ?", likeProduct)}
		sq, args, err := squirrel.Select("id").Distinct().From("affected").Where(cond).ToSql()
		if err != nil {
			return ds, err
		}
		bldr = bldr.Join(fmt.Sprintf("(%s) affect on affect.id = vulnview.id", sq), args...)
	}
	if len(cve) > 0 {
		sq, args, err := squirrel.Select("id").Distinct().From("related").Where(squirrel.Eq{"vulinfo_id": cve}).ToSql()
		if err != nil {
			return ds, err
		}
		bldr = bldr.Join(fmt.Sprintf("(%s) relate on relate.id = vulnview.id", sq), args...)
	}

	//query
	if psql, args, err := bldr.OrderBy("date_update desc", "id").ToSql(); err != nil {
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
