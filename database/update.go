package database

import (
	"time"

	"github.com/spiegel-im-spiegel/go-myjvn/rss"
	"github.com/spiegel-im-spiegel/go-myjvn/rss/option"
	"github.com/spiegel-im-spiegel/go-myjvn/status"
	"github.com/spiegel-im-spiegel/go-myjvn/vuldef"
)

//Update returns result for updating JVN data
func (db *DB) Update(month bool) ([]string, error) {
	ids := []string{}
	if db == nil {
		return ids, nil
	}

	start, err := db.getLastUpdate()
	if err != nil {
		return ids, err
	}
	jvnrss, err := db.getJVNRSS(start, month)
	if err != nil {
		return ids, err
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return ids, err
	}

	//prepared statements
	addStmt, err := tx.Prepare("insert into vulnlist (id, title, description, uri, creator, date_public, date_publish, date_update) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return ids, err
	}
	defer addStmt.Close()
	updStmt, err := tx.Prepare("update vulnlist set title = ?, description = ?, uri = ?, creator = ?, date_public = ?, date_publish = ?, date_update = ? where id = ?")
	if err != nil {
		tx.Rollback()
		return ids, err
	}
	defer updStmt.Close()

	//insert and update data
	err = nil
	for _, itm := range jvnrss.Items {
		upd := db.getModifiedDate(itm.Identifier)
		if upd.IsZero() {
			ids = append(ids, itm.Identifier)
			if _, err = addStmt.Exec(itm.Identifier, itm.Title, itm.Description, itm.Link, itm.Creator, itm.Date.Unix(), itm.Issued.Unix(), itm.Modified.Unix()); err != nil {
				break
			}
		} else if itm.Modified.After(upd) {
			ids = append(ids, itm.Identifier)
			if _, err = updStmt.Exec(itm.Title, itm.Description, itm.Link, itm.Creator, itm.Date.Unix(), itm.Issued.Unix(), itm.Modified.Unix(), itm.Identifier); err != nil {
				break
			}
		}
	}

	if err != nil {
		tx.Rollback()
		return ids, err
	}
	return ids, tx.Commit()
}

//UpdateDetail returns result for updating JVN detail data
func (db *DB) UpdateDetail(ids []string) error {
	if db == nil {
		return nil
	}
	if len(ids) == 0 {
		return nil
	}

	vulnInfo, err := db.getVULDEF(ids)
	if err != nil {
		return err
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	//prepared statements
	updList, err := tx.Prepare("update vulnlist set impact = ?, solution = ?, date_public = ?, date_publish = ? where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer updList.Close()
	delAffected, err := tx.Prepare("delete from affected where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer delAffected.Close()
	addAffected, err := tx.Prepare("insert into affected (id, name, product_name, version_number) values(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer addAffected.Close()
	delCVSS, err := tx.Prepare("delete from cvss where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer delCVSS.Close()
	addCVSS, err := tx.Prepare("insert into cvss (id, version, base_vector, base_score, severity) values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer addCVSS.Close()
	delRelated, err := tx.Prepare("delete from related where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer delRelated.Close()
	addRelated, err := tx.Prepare("insert into related (id, type, name, vulinfo_id, title, url) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer addRelated.Close()
	delHistory, err := tx.Prepare("delete from history where id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer delHistory.Close()
	addHistory, err := tx.Prepare("insert into history (id, history_no, description, date_time) values(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer addHistory.Close()

	//update detail data
	err = nil
	for _, vuln := range vulnInfo.Vulinfo {
		//update vulnlist
		if _, err = updList.Exec(vuln.VulinfoData.Impact.Description, vuln.VulinfoData.Solution.Description, vuln.VulinfoData.DatePublic.Unix(), vuln.VulinfoData.DateFirstPublished.Unix(), vuln.VulinfoID); err != nil {
			break
		}
		//Affected info.
		if _, err = delAffected.Exec(vuln.VulinfoID); err != nil {
			break
		}
		for _, affected := range vuln.VulinfoData.Affected {
			for _, ver := range affected.VersionNumber {
				if _, err = addAffected.Exec(vuln.VulinfoID, affected.Name, affected.ProductName, ver); err != nil {
					break
				}
			}
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}

		//Impact info
		if _, err = delCVSS.Exec(vuln.VulinfoID); err != nil {
			break
		}
		for _, cvss := range vuln.VulinfoData.Impact.CVSS {
			if _, err = addCVSS.Exec(vuln.VulinfoID, cvss.Version, cvss.BaseVector, cvss.BaseScore, cvss.Severity); err != nil {
				break
			}
		}
		if err != nil {
			break
		}

		//Related info.
		if _, err = delRelated.Exec(vuln.VulinfoID); err != nil {
			break
		}
		for _, related := range vuln.VulinfoData.Related {
			if _, err = addRelated.Exec(vuln.VulinfoID, related.Type, related.Name, related.VulinfoID, related.Title, related.URL); err != nil {
				break
			}
		}
		if err != nil {
			break
		}

		//History info.
		if _, err = delHistory.Exec(vuln.VulinfoID); err != nil {
			break
		}
		for _, history := range vuln.VulinfoData.History {
			if _, err = addHistory.Exec(vuln.VulinfoID, history.HistoryNo, history.Description, history.DateTime.Unix()); err != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (db *DB) getLastUpdate() (time.Time, error) {
	t := time.Time{}
	rows, err := db.GetDB().Query("select max(date_update) as last from vulnlist")
	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var utime int64
		rows.Scan(&utime)
		if utime > 0 {
			t = time.Unix(utime, 0)
		}
	}
	return t, rows.Err()
}

func (db *DB) getModifiedDate(id string) time.Time {
	var utime int64
	if err := db.GetDB().QueryRow("select date_update from vulnlist where id = ?", id).Scan(&utime); err != nil {
		return time.Time{}
	}
	if utime > 0 {
		return time.Unix(utime, 0)
	}
	return time.Time{}
}

func (db *DB) getJVNRSS(start time.Time, month bool) (*rss.JVNRSS, error) {
	startItem := 1
	maxItem := 0
	mode := option.RangeWeek
	if month {
		mode = option.RangeMonth
	}
	jvnrss := (*rss.JVNRSS)(nil)
	for {
		var opt *option.Option
		if start.IsZero() {
			opt = option.New(
				option.WithRangeDatePublicMode(option.NoRange),
				option.WithRangeDatePublishedMode(mode),
				option.WithRangeDateFirstPublishedMode(option.NoRange),
				option.WithStartItem(startItem),
			)
		} else {
			opt = option.New(
				option.WithRangeDatePublicMode(option.NoRange),
				option.WithRangeDatePublishedPeriod(start, time.Now()),
				option.WithRangeDateFirstPublishedMode(option.NoRange),
				option.WithStartItem(startItem),
			)
		}
		if startItem == 1 {
			xml, err := db.api.VulnOverviewListXML(opt)
			if err != nil {
				return jvnrss, err
			}
			stat, err := status.Unmarshal(xml)
			if err != nil {
				return jvnrss, err
			}
			if err2 := stat.GetError(); err2 != nil {
				return jvnrss, err2
			}
			maxItem = stat.Status.TotalRes
			r, err := rss.Unmarshal(xml)
			if err != nil {
				return jvnrss, err
			}
			startItem += len(r.Items)
			jvnrss = r
		} else {
			r, err := db.api.VulnOverviewList(opt)
			if err != nil {
				return jvnrss, err
			}
			startItem += len(r.Items)
			jvnrss.Append(r)
		}
		if maxItem < startItem {
			break
		}
	}
	return jvnrss, nil
}

func (db *DB) getVULDEF(idlist []string) (*vuldef.VULDEF, error) {
	vulnInfo := (*vuldef.VULDEF)(nil)
	for i := 0; i < len(idlist); i += vuldef.MaxItems {
		var ids []string
		if i+vuldef.MaxItems < len(idlist) {
			ids = idlist[i : i+vuldef.MaxItems]
		} else {
			ids = idlist[i:]
		}
		vuln, err := db.api.VulnDetailInfo(ids)
		if err != nil {
			return vulnInfo, err
		}
		if vulnInfo == nil {
			vulnInfo = vuln
		} else {
			vulnInfo.Append(vuln)
		}
	}
	return vulnInfo, nil
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
