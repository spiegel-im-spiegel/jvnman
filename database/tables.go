package database

import (
	"database/sql"
	"time"

	squirrel "gopkg.in/Masterminds/squirrel.v1"
)

//SQL statements
var (
	deleteAffected = squirrel.Delete("affected")
	deleteCVSS     = squirrel.Delete("cvss")
	deleteRelated  = squirrel.Delete("related")
	deleteHistory  = squirrel.Delete("history")
	selectVulnview = squirrel.Select(
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
	selectAffected = squirrel.Select(
		"id",
		"name",
		"product_name",
		"version_number",
	).From("affected").OrderBy("name", "product_name", "version_number")
	selectCVSS = squirrel.Select(
		"id",
		"version",
		"base_vector",
		"base_score",
		"severity",
	).From("cvss").Where(squirrel.Eq{"version": "3.0"})
	selectRelated = squirrel.Select(
		"id",
		"type",
		"name",
		"vulinfo_id",
		"title",
		"url",
	).From("related").OrderBy("type", "vulinfo_id", "name", "title", "url")
)

//Vulnlist is definition of vulnlist table
type Vulnlist struct {
	ID          sql.NullString `db:"id,primarykey"`
	Title       sql.NullString `db:"title"`
	Description sql.NullString `db:"description"`
	URI         sql.NullString `db:"uri"`
	Creator     sql.NullString `db:"creator"`
	Impact      sql.NullString `db:"impact"`
	Solution    sql.NullString `db:"solution"`
	DatePublic  sql.NullInt64  `db:"date_public"`
	DatePublish sql.NullInt64  `db:"date_publish"`
	DateUpdate  sql.NullInt64  `db:"date_update"`
}

//NewVulnlist returns Vulnlist instance
func NewVulnlist(id, title, description, uri, creator, impact, solution string, datePublic, datePublish, dateUpdate int64) *Vulnlist {
	return &Vulnlist{
		ID:          Text(id),
		Title:       Text(title),
		Description: Text(description),
		URI:         Text(uri),
		Creator:     Text(creator),
		Impact:      Text(impact),
		Solution:    Text(solution),
		DatePublic:  Integer(datePublic),
		DatePublish: Integer(datePublish),
		DateUpdate:  Integer(dateUpdate),
	}
}

//GetDatePublic returns time.Time instance of Vulnlist.DatePublic
func (ds *Vulnlist) GetDatePublic() time.Time {
	if ds.DatePublic.Valid {
		return getTimeFromUnixtime(ds.DatePublic.Int64)
	}
	return time.Time{}
}

//GetDatePublish returns time.Time instance of Vulnlist.DatePublish
func (ds *Vulnlist) GetDatePublish() time.Time {
	if ds.DatePublish.Valid {
		return getTimeFromUnixtime(ds.DatePublish.Int64)
	}
	return time.Time{}
}

//GetDateUpdate returns time.Time instance of Vulnlist.DateUpdate
func (ds *Vulnlist) GetDateUpdate() time.Time {
	if ds.DateUpdate.Valid {
		return getTimeFromUnixtime(ds.DateUpdate.Int64)
	}
	return time.Time{}
}

//Affected is definition of affected table
type Affected struct {
	ID            sql.NullString `db:"id,primarykey"`
	Name          sql.NullString `db:"name,primarykey"`
	ProductName   sql.NullString `db:"product_name,primarykey"`
	VersionNumber sql.NullString `db:"version_number,primarykey"`
}

//NewAffected returns Affected instance
func NewAffected(id, name, productName, versionNumber string) *Affected {
	return &Affected{
		ID:            Text(id),
		Name:          Text(name),
		ProductName:   Text(productName),
		VersionNumber: Text(versionNumber),
	}
}

//CVSS is definition of cvss table
type CVSS struct {
	ID         sql.NullString  `db:"id,primarykey"`
	Version    sql.NullString  `db:"version,primarykey"`
	BaseVector sql.NullString  `db:"base_vector"`
	BaseScore  sql.NullFloat64 `db:"base_score"`
	Severity   sql.NullString  `db:"severity"`
}

//NewCVSS returns CVSS instance
func NewCVSS(id, version, baseVector, severity string, baseScore float64) *CVSS {
	return &CVSS{
		ID:         Text(id),
		Version:    Text(version),
		BaseVector: Text(baseVector),
		BaseScore:  Numeric(baseScore),
		Severity:   Text(severity),
	}
}

//Related is definition of related table
type Related struct {
	ID        sql.NullString `db:"id,primarykey"`
	Type      sql.NullString `db:"type,primarykey"`
	Name      sql.NullString `db:"name,primarykey"`
	VulinfoID sql.NullString `db:"vulinfo_id,primarykey"`
	Title     sql.NullString `db:"title"`
	URL       sql.NullString `db:"url"`
}

//NewRelated returns Related instance
func NewRelated(id, typeR, name, vulinfoID, title, url string) *Related {
	return &Related{
		ID:        Text(id),
		Type:      Text(typeR),
		Name:      Text(name),
		VulinfoID: Text(vulinfoID),
		Title:     Text(title),
		URL:       Text(url),
	}

}

//History is definition of history table
type History struct {
	ID          sql.NullString `db:"id,primarykey"`
	HistoryNo   sql.NullInt64  `db:"history_no,primarykey"`
	Description sql.NullString `db:"description"`
	DateTime    sql.NullInt64  `db:"date_time"`
}

//NewHistory returns History instance
func NewHistory(id string, historyNo int64, description string, dateTime int64) *History {
	return &History{
		ID:          Text(id),
		HistoryNo:   Integer(historyNo),
		Description: Text(description),
		DateTime:    Integer(dateTime),
	}
}

//GetDateTime returns time.Time instance of History.DateTime
func (ds *History) GetDateTime() time.Time {
	if ds.DateTime.Valid {
		return getTimeFromUnixtime(ds.DateTime.Int64)
	}
	return time.Time{}
}

//Vulnview is definition of vulnview view
type Vulnview struct {
	ID           sql.NullString  `db:"id"`
	Title        sql.NullString  `db:"title"`
	Description  sql.NullString  `db:"description"`
	URI          sql.NullString  `db:"uri"`
	Impact       sql.NullString  `db:"impact"`
	Solution     sql.NullString  `db:"solution"`
	CVSSScore    sql.NullFloat64 `db:"cvss_score"`
	CVSSSeverity sql.NullString  `db:"cvss_severity"`
	DatePublic   sql.NullInt64   `db:"date_public"`
	DatePublish  sql.NullInt64   `db:"date_publish"`
	DateUpdate   sql.NullInt64   `db:"date_update"`
}

//GetDatePublic returns time.Time instance of Vulnview.DatePublic
func (ds *Vulnview) GetDatePublic() time.Time {
	if ds.DatePublic.Valid {
		return getTimeFromUnixtime(ds.DatePublic.Int64)
	}
	return time.Time{}
}

//GetDatePublish returns time.Time instance of Vulnview.DatePublish
func (ds *Vulnview) GetDatePublish() time.Time {
	if ds.DatePublish.Valid {
		return getTimeFromUnixtime(ds.DatePublish.Int64)
	}
	return time.Time{}
}

//GetDateUpdate returns time.Time instance of Vulnview.DateUpdate
func (ds *Vulnview) GetDateUpdate() time.Time {
	if ds.DateUpdate.Valid {
		return getTimeFromUnixtime(ds.DateUpdate.Int64)
	}
	return time.Time{}
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
