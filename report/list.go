package report

import (
	"bytes"
	"encoding/csv"
	"fmt"
	html "html/template"
	"io"
	"strings"
	text "text/template"

	"github.com/spiegel-im-spiegel/jvnman/database"
)

//VulnInfo is dataset for report
type VulnInfo struct {
	ID          string
	Title       string
	Description string
	URI         string
	Impact      string
	Solution    string
	Severity    string
	DatePublic  string
	DatePublish string
	DateUpdate  string
}

var csvHeader = []string{
	"ID",
	"タイトル",
	"概要",
	"URI",
	"想定される影響",
	"対策",
	"深刻度",
	"発見日",
	"公開日",
	"最終更新日",
}

//ListData returns io.Reader for listing
func ListData(db *database.DB, days int, score float64, product, cve string, f Format, verbose bool) (io.Reader, error) {
	view, err := db.GetVulnviewList(days, score, product, cve)
	if err != nil {
		return nil, err
	}
	list := []VulnInfo{}
	for _, v := range view {
		l := VulnInfo{
			ID:          v.ID.String,
			Title:       v.Title.String,
			Description: v.Description.String,
			URI:         v.URI.String,
			Impact:      v.Impact.String,
			Solution:    v.Solution.String,
			Severity:    fmt.Sprintf("%v (%.1f)", getSeverityJa(v.CVSSSeverity.String), v.CVSSScore.Float64),
			DatePublic:  v.GetDatePublic().Format("2006年1月2日"),
			DatePublish: v.GetDatePublish().Format("2006年1月2日"),
			DateUpdate:  v.GetDateUpdate().Format("2006年1月2日"),
		}
		list = append(list, l)
	}

	buf := &bytes.Buffer{}
	switch f {
	case FormHTML:
		var fname string
		if verbose {
			fname = "/template-list-detail.html"
		} else {
			fname = "/template-list.html"
		}
		tf, errAssets := Assets.Open(fname)
		if errAssets != nil {
			return nil, errAssets
		}
		tmpdata := &strings.Builder{}
		io.Copy(tmpdata, tf)
		t, errTmp := html.New("Listing by HTML").Parse(tmpdata.String())
		if errTmp != nil {
			return nil, errTmp
		}
		if err = t.Execute(buf, list); err != nil {
			return nil, err
		}
	case FormMarkdown:
		var fname string
		if verbose {
			fname = "/template-list-detail.md"
		} else {
			fname = "/template-list.md"
		}
		tf, errAssets := Assets.Open(fname)
		if errAssets != nil {
			return nil, errAssets
		}
		tmpdata := &strings.Builder{}
		io.Copy(tmpdata, tf)
		t, errTmp := text.New("Listing by Markdown").Parse(tmpdata.String())
		if errTmp != nil {
			return nil, errTmp
		}
		if err = t.Execute(buf, list); err != nil {
			return nil, err
		}
	case FormCSV:
		w := csv.NewWriter(buf)
		w.Write(csvHeader)
		for _, l := range list {
			rec := []string{
				l.ID,
				l.Title,
				l.Description,
				l.URI,
				l.Impact,
				l.Solution,
				l.Severity,
				l.DatePublic,
				l.DatePublish,
				l.DateUpdate,
			}
			w.Write(rec)
		}
		w.Flush()
	default:
	}
	return buf, nil
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
