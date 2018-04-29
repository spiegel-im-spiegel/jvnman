package database

import (
	"time"
)

//Vulnlist is definition of vulnlist table
type Vulnlist struct {
	ID          string `db:"id,primarykey"`
	Title       string `db:"title"`
	Description string `db:"description"`
	URI         string `db:"uri"`
	Creator     string `db:"creator"`
	Impact      string `db:"impact"`
	Solution    string `db:"solution"`
	DatePublic  int64  `db:"date_public"`
	DatePublish int64  `db:"date_publish"`
	DateUpdate  int64  `db:"date_update"`
}

//NewVulnlist returns Vulnlist instance
func NewVulnlist(id, title, description, uri, creator, impact, solution string, datePublic, datePublish, dateUpdate int64) *Vulnlist {
	return &Vulnlist{
		ID:          id,
		Title:       title,
		Description: description,
		URI:         uri,
		Creator:     creator,
		Impact:      impact,
		Solution:    solution,
		DatePublic:  datePublic,
		DatePublish: datePublish,
		DateUpdate:  dateUpdate,
	}
}

//GetDatePublic returns time.Time instance of Vulnlist.DatePublic
func (ds *Vulnlist) GetDatePublic() time.Time {
	return getTimeFromUnixtime(ds.DatePublic)
}

//GetDatePublish returns time.Time instance of Vulnlist.DatePublish
func (ds *Vulnlist) GetDatePublish() time.Time {
	return getTimeFromUnixtime(ds.DatePublish)
}

//GetDateUpdate returns time.Time instance of Vulnlist.DateUpdate
func (ds *Vulnlist) GetDateUpdate() time.Time {
	return getTimeFromUnixtime(ds.DateUpdate)
}

//Affected is definition of affected table
type Affected struct {
	ID            string `db:"id,primarykey"`
	Name          string `db:"name,primarykey"`
	ProductName   string `db:"product_name,primarykey"`
	VersionNumber string `db:"version_number,primarykey"`
}

//NewAffected returns Affected instance
func NewAffected(id, name, productName, versionNumber string) *Affected {
	return &Affected{
		ID:            id,
		Name:          name,
		ProductName:   productName,
		VersionNumber: versionNumber,
	}
}

//CVSS is definition of cvss table
type CVSS struct {
	ID         string  `db:"id,primarykey"`
	Version    string  `db:"version,primarykey"`
	BaseVector string  `db:"base_vector"`
	BaseScore  float64 `db:"base_score"`
	Severity   string  `db:"severity"`
}

//NewCVSS returns CVSS instance
func NewCVSS(id, version, baseVector, severity string, baseScore float64) *CVSS {
	return &CVSS{
		ID:         id,
		Version:    version,
		BaseVector: baseVector,
		BaseScore:  baseScore,
		Severity:   severity,
	}
}

//Related is definition of related table
type Related struct {
	ID        string `db:"id,primarykey"`
	Type      string `db:"type,primarykey"`
	Name      string `db:"name,primarykey"`
	VulinfoID string `db:"vulinfo_id,primarykey"`
	Title     string `db:"title"`
	URL       string `db:"url"`
}

//NewRelated returns Related instance
func NewRelated(id, typeR, name, vulinfoID, title, url string) *Related {
	return &Related{
		ID:        id,
		Type:      typeR,
		Name:      name,
		VulinfoID: vulinfoID,
		Title:     title,
		URL:       url,
	}

}

//History is definition of history table
type History struct {
	ID          string `db:"id,primarykey"`
	HistoryNo   int    `db:"history_no,primarykey"`
	Description string `db:"description"`
	DateTime    int64  `db:"date_time"`
}

//NewHistory returns History instance
func NewHistory(id string, historyNo int, description string, dateTime int64) *History {
	return &History{
		ID:          id,
		HistoryNo:   historyNo,
		Description: description,
		DateTime:    dateTime,
	}
}

//GetDateTime returns time.Time instance of History.DateTime
func (ds *History) GetDateTime() time.Time {
	return getTimeFromUnixtime(ds.DateTime)
}

//Vulnview is definition of vulnview view
type Vulnview struct {
	ID           string  `db:"id"`
	Title        string  `db:"title"`
	Description  string  `db:"description"`
	URI          string  `db:"uri"`
	Impact       string  `db:"impact"`
	Solution     string  `db:"solution"`
	CVSSScore    float64 `db:"cvss_score"`
	CVSSSeverity string  `db:"cvss_severity"`
	DatePublic   int64   `db:"date_public"`
	DatePublish  int64   `db:"date_publish"`
	DateUpdate   int64   `db:"date_update"`
}

//GetDatePublic returns time.Time instance of Vulnview.DatePublic
func (ds *Vulnview) GetDatePublic() time.Time {
	return getTimeFromUnixtime(ds.DatePublic)
}

//GetDatePublish returns time.Time instance of Vulnview.DatePublish
func (ds *Vulnview) GetDatePublish() time.Time {
	return getTimeFromUnixtime(ds.DatePublish)
}

//GetDateUpdate returns time.Time instance of Vulnview.DateUpdate
func (ds *Vulnview) GetDateUpdate() time.Time {
	return getTimeFromUnixtime(ds.DateUpdate)
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
