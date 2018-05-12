package report

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/spiegel-im-spiegel/jvnman/database"
)

var csvHeader = []string{
	"ID",
	"タイトル",
	"概要",
	"URI",
	"想定される影響",
	"対策",
	"CVSSv3ベクタ",
	"深刻度",
	"発見日",
	"公開日",
	"最終更新日",
}

//ListData returns io.Reader for listing
func ListData(db *database.DB, days int, score float64, product, cve string, f Format, tfname string) (io.Reader, error) {
	buf := &bytes.Buffer{}
	view, err := db.GetVulnviewList(days, score, product, cve)
	if err != nil {
		return buf, err
	}
	list := []VulnInfo{}
	for _, v := range view {
		severity := ""
		if len(v.CVSSSeverity.String) > 0 {
			severity = fmt.Sprintf("%v (%.1f)", getSeverityJa(v.CVSSSeverity.String), v.CVSSScore.Float64)
		}
		l := VulnInfo{
			ID:          v.ID.String,
			Title:       v.Title.String,
			Description: v.Description.String,
			URI:         v.URI.String,
			Impact:      v.Impact.String,
			Solution:    v.Solution.String,
			CVSSVector:  v.CVSSVector.String,
			Severity:    severity,
			DatePublic:  v.GetDatePublic().Format("2006年1月2日"),
			DatePublish: v.GetDatePublish().Format("2006年1月2日"),
			DateUpdate:  v.GetDateUpdate().Format("2006年1月2日"),
		}
		list = append(list, l)
	}

	if f == FormCSV {
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
				l.CVSSVector,
				l.Severity,
				l.DatePublic,
				l.DatePublish,
				l.DateUpdate,
			}
			w.Write(rec)
		}
		w.Flush()
	} else {
		var tf io.Reader
		if len(tfname) > 0 {
			file, err := os.Open(tfname)
			if err != nil {
				return buf, err
			}
			tf = file
		} else {
			file, err := Assets.Open("/template-list-detail.md")
			if err != nil {
				return buf, err
			}
			tf = file
		}
		tmpdata := &strings.Builder{}
		io.Copy(tmpdata, tf)
		t, err := template.New("Listing by Markdown").Parse(tmpdata.String())
		if err != nil {
			return buf, err
		}
		if err := t.Execute(buf, list); err != nil {
			return buf, err
		}
		if f == FormHTML {
			return convHTML(buf), nil
		}
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
