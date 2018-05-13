package database

import (
	"database/sql"
	"html"
	"time"

	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"gopkg.in/Masterminds/squirrel.v1"
	gorp "gopkg.in/gorp.v2"
)

//Update update JVN vulnerability data
func (db *DB) Update(month bool, keyword string) error {
	if db == nil {
		return nil
	}
	db.GetLogger().Debugln("month option:", month)

	jvnrss, err := db.GetJVNRSS(db.GetLastUpdate(), month, keyword)
	if err != nil {
		return err
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	if err := db.update(tx, jvnrss); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

//UpdateID update JVN vulnerability data by ID
func (db *DB) UpdateID(id string) error {
	if db == nil {
		return nil
	}
	db.GetLogger().Debugln("update id:", id)

	jvnrss, err := db.GetJVNRSSByKeyword(id)
	if err != nil {
		return err
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	if err := db.update(tx, jvnrss); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

//Update returns result for updating JVN data
func (db *DB) update(tx *gorp.Transaction, jvnrss *rss.JVNRSS) error {
	ids := []string{}
	for _, itm := range jvnrss.Items {
		obj, err := tx.Get(Vulnlist{}, itm.Identifier)
		if err != nil {
			return err
		}
		if obj == nil {
			db.GetLogger().Debugln("Insert", itm.Identifier)
			if err := tx.Insert(NewVulnlist(itm.Identifier, html.UnescapeString(itm.Title), html.UnescapeString(itm.Description), html.UnescapeString(itm.Link), html.UnescapeString(itm.Creator), "", "", itm.Date.Unix(), itm.Issued.Unix(), itm.Modified.Unix())); err != nil {
				return err
			}
			ids = append(ids, itm.Identifier)
		} else if ds, ok := obj.(*Vulnlist); ok && itm.Modified.After(ds.GetDateUpdate()) {
			db.GetLogger().Debugln("Update", itm.Identifier)
			ds.Title = Text(html.UnescapeString(itm.Title))
			ds.Description = Text(html.UnescapeString(itm.Description))
			ds.URI = Text(html.UnescapeString(itm.Link))
			ds.Creator = Text(html.UnescapeString(itm.Creator))
			ds.DatePublic = Integer(itm.Date.Unix())
			ds.DatePublish = Integer(itm.Issued.Unix())
			ds.DateUpdate = Integer(itm.Modified.Unix())
			if _, err := tx.Update(ds); err != nil {
				return err
			}
			ids = append(ids, itm.Identifier)
		}
	}
	return db.updateDetail(tx, ids)
}

func (db *DB) updateDetail(tx *gorp.Transaction, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	db.GetLogger().Debugln("JVN Vulnerability ID:", ids)

	vulnInfo, err := db.GetVULDEF(ids)
	if err != nil {
		return err
	}

	//update detail data
	for _, vuln := range vulnInfo.Vulinfo {
		//update vulnlist
		if obj, err := tx.Get(Vulnlist{}, vuln.VulinfoID); err != nil {
			return err
		} else if obj != nil {
			ds := obj.(*Vulnlist)
			ds.Impact = Text(html.UnescapeString(vuln.VulinfoData.Impact.Description))
			ds.Solution = Text(html.UnescapeString(vuln.VulinfoData.Solution.Description))
			ds.DatePublic = Integer(vuln.VulinfoData.DatePublic.Unix())
			ds.DatePublish = Integer(vuln.VulinfoData.DateFirstPublished.Unix())
			if _, err := tx.Update(ds); err != nil {
				return err
			}
		}

		//Affected info.
		if psql, args, err := deleteAffected.Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			return err
		} else if _, err = tx.Exec(psql, args...); err != nil {
			return err
		}
		for _, affected := range vuln.VulinfoData.Affected {
			for _, ver := range affected.VersionNumber {
				if err := tx.Insert(NewAffected(vuln.VulinfoID, html.UnescapeString(affected.Name), html.UnescapeString(affected.ProductName), html.UnescapeString(ver))); err != nil {
					return err
				}
			}
		}

		//Impact info
		if psql, args, err := deleteCVSS.Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			return err
		} else if _, err = tx.Exec(psql, args...); err != nil {
			return err
		}
		for _, cvss := range vuln.VulinfoData.Impact.CVSS {
			if err := tx.Insert(NewCVSS(vuln.VulinfoID, cvss.Version, cvss.BaseVector, cvss.Severity, getFloatFromString(cvss.BaseScore))); err != nil {
				return err
			}
		}
		//Related info.
		if psql, args, err := deleteRelated.Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			return err
		} else if _, err = tx.Exec(psql, args...); err != nil {
			return err
		}
		for _, related := range vuln.VulinfoData.Related {
			if err := tx.Insert(NewRelated(vuln.VulinfoID, related.Type, html.UnescapeString(related.Name), html.UnescapeString(related.VulinfoID), html.UnescapeString(related.Title), html.UnescapeString(related.URL))); err != nil {
				return err
			}
		}

		//History info.
		if psql, args, err := deleteHistory.Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			return err
		} else if _, err = tx.Exec(psql, args...); err != nil {
			return err
		}
		for _, history := range vuln.VulinfoData.History {
			if err := tx.Insert(NewHistory(vuln.VulinfoID, int64(history.HistoryNo), html.UnescapeString(history.Description), history.DateTime.Unix())); err != nil {
				return err
			}
		}
	}
	return nil
}

//GetLastUpdate returns  last update time.Time
func (db *DB) GetLastUpdate() time.Time {
	var ds struct {
		Last sql.NullInt64 `db:"last"`
	}
	if psql, _, err := squirrel.Select("max(date_update) as last").From("vulnlist").ToSql(); err != nil {
		db.GetLogger().Errorln(err)
		return time.Time{}
	} else if err := db.GetDB().SelectOne(&ds, psql); err != nil {
		db.GetLogger().Errorln(err)
		return time.Time{}
	}
	if !ds.Last.Valid {
		db.GetLogger().Println("no data in database")
		return time.Time{}
	}
	dt := getTimeFromUnixtime(ds.Last.Int64)
	db.GetLogger().Println("last update:", dt)
	return dt
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
