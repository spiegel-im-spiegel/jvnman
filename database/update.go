package database

import (
	"html"
	"time"

	"github.com/Masterminds/squirrel"
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
	logger := db.GetLogger()
	logger.Println("Update JVN data:", db.GetDBFile())

	jvnrss, err := db.getJVNRSS(db.getLastUpdate(), month)
	if err != nil {
		return ids, err
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return ids, err
	}

	//insert and update data
	err = nil
	for _, itm := range jvnrss.Items {
		obj, errGet := db.GetDB().Get(Vulnlist{}, itm.Identifier)
		if errGet != nil {
			err = errGet
			break
		}
		if obj == nil {
			ids = append(ids, itm.Identifier)
			if err = tx.Insert(NewVulnlist(itm.Identifier, html.UnescapeString(itm.Title), html.UnescapeString(itm.Description), html.UnescapeString(itm.Link), html.UnescapeString(itm.Creator), "", "", itm.Date.Unix(), itm.Issued.Unix(), itm.Modified.Unix())); err != nil {
				break
			}
		} else {
			ds := obj.(*Vulnlist)
			if itm.Modified.After(ds.GetDateUpdate()) {
				ds.Title = Text(html.UnescapeString(itm.Title))
				ds.Description = Text(html.UnescapeString(itm.Description))
				ds.URI = Text(html.UnescapeString(itm.Link))
				ds.Creator = Text(html.UnescapeString(itm.Creator))
				ds.DatePublic = Integer(itm.Date.Unix())
				ds.DatePublish = Integer(itm.Issued.Unix())
				ds.DateUpdate = Integer(itm.Modified.Unix())
				if _, err = tx.Update(ds); err != nil {
					break
				}
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
	logger := db.GetLogger()
	logger.Println("Update JVN data detail:", db.GetDBFile())
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

	//update detail data
	err = nil
	for _, vuln := range vulnInfo.Vulinfo {
		//update vulnlist
		obj, errGet := db.GetDB().Get(Vulnlist{}, vuln.VulinfoID)
		if errGet != nil {
			err = errGet
			break
		}
		if obj != nil {
			ds := obj.(*Vulnlist)
			ds.Impact = Text(html.UnescapeString(vuln.VulinfoData.Impact.Description))
			ds.Solution = Text(html.UnescapeString(vuln.VulinfoData.Solution.Description))
			ds.DatePublic = Integer(vuln.VulinfoData.DatePublic.Unix())
			ds.DatePublish = Integer(vuln.VulinfoData.DateFirstPublished.Unix())
			if _, err = tx.Update(ds); err != nil {
				break
			}
		}

		//Affected info.
		if psql, args, err := squirrel.Delete("affected").Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			break
		} else if _, err = tx.Exec(psql, args...); err != nil {
			break
		}
		for _, affected := range vuln.VulinfoData.Affected {
			for _, ver := range affected.VersionNumber {
				if err = tx.Insert(NewAffected(vuln.VulinfoID, html.UnescapeString(affected.Name), html.UnescapeString(affected.ProductName), html.UnescapeString(ver))); err != nil {
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
		if psql, args, err := squirrel.Delete("cvss").Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			break
		} else if _, err = tx.Exec(psql, args...); err != nil {
			break
		}
		for _, cvss := range vuln.VulinfoData.Impact.CVSS {
			if err = tx.Insert(NewCVSS(vuln.VulinfoID, cvss.Version, cvss.BaseVector, cvss.Severity, getFloatFromString(cvss.BaseScore))); err != nil {
				break
			}
		}
		if err != nil {
			break
		}

		//Related info.
		if psql, args, err := squirrel.Delete("related").Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			break
		} else if _, err = tx.Exec(psql, args...); err != nil {
			break
		}
		for _, related := range vuln.VulinfoData.Related {
			if err = tx.Insert(NewRelated(vuln.VulinfoID, related.Type, html.UnescapeString(related.Name), html.UnescapeString(related.VulinfoID), html.UnescapeString(related.Title), html.UnescapeString(related.URL))); err != nil {
				break
			}
		}
		if err != nil {
			break
		}

		//History info.
		if psql, args, err := squirrel.Delete("history").Where(squirrel.Eq{"id": vuln.VulinfoID}).ToSql(); err != nil {
			break
		} else if _, err = tx.Exec(psql, args...); err != nil {
			break
		}
		for _, history := range vuln.VulinfoData.History {
			if err = tx.Insert(NewHistory(vuln.VulinfoID, int64(history.HistoryNo), html.UnescapeString(history.Description), history.DateTime.Unix())); err != nil {
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

func (db *DB) getLastUpdate() time.Time {
	var ds struct {
		utime int64
	}
	if err := db.GetDB().SelectOne(&ds, "select max(date_update) as last from vulnlist"); err != nil {
		return time.Time{}
	}
	return getTimeFromUnixtime(ds.utime)
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
			xml, err := db.GetAPI().VulnOverviewListXML(opt)
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
			r, err := db.GetAPI().VulnOverviewList(opt)
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
		vuln, err := db.GetAPI().VulnDetailInfo(ids)
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
