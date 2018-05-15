package report

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/spiegel-im-spiegel/go-cvss/v3/base"
	"github.com/spiegel-im-spiegel/jvnman/database"
	"golang.org/x/text/language"
)

//Info returns io.Reader for detail info
func Info(db *database.DB, id string, tf io.Reader, f Format) (io.Reader, error) {
	buf := &bytes.Buffer{}
	detail := VulnDetail{}
	v := getGetVulnviewByID(db, id)
	if v == nil {
		return buf, nil
	}
	severity := ""
	if len(v.CVSSSeverity.String) > 0 {
		severity = fmt.Sprintf("%v (%.1f)", getSeverityJa(v.CVSSSeverity.String), v.CVSSScore.Float64)
	}
	detail.Info.ID = v.ID.String
	detail.Info.Title = v.Title.String
	detail.Info.Description = v.Description.String
	detail.Info.URI = v.URI.String
	detail.Info.Impact = v.Impact.String
	detail.Info.Solution = v.Solution.String
	detail.Info.CVSSVector = v.CVSSVector.String
	detail.Info.Severity = severity
	detail.Info.DatePublic = v.GetDatePublic().Format("2006年1月2日")
	detail.Info.DatePublish = v.GetDatePublish().Format("2006年1月2日")
	detail.Info.DateUpdate = v.GetDateUpdate().Format("2006年1月2日")

	af := []AffectInfo{}
	dsA := db.GetAffected(id)
	if dsA != nil {
		for _, a := range dsA {
			aa := AffectInfo{
				ID:            a.ID.String,
				Name:          a.Name.String,
				ProductName:   a.ProductName.String,
				VersionNumber: a.VersionNumber.String,
			}
			af = append(af, aa)
		}
	}
	detail.Affects = af
	c := db.GetCVSS(id)
	if c != nil {
		detail.CVSS.ID = c.ID.String
		detail.CVSS.Version = c.Version.String
		detail.CVSS.BaseVector = c.BaseVector.String
		if m, err := base.Decode(detail.CVSS.BaseVector); err == nil {
			file, err := Assets.Open("/cvss.md")
			if err != nil {
				return buf, err
			}
			defer file.Close()
			if rep, err := m.Report(file, language.Japanese); err == nil {
				bldr := &strings.Builder{}
				io.Copy(bldr, rep)
				detail.CVSS.BaseReport = bldr.String()
			}
		}
		detail.CVSS.BaseScore = c.BaseScore.Float64
		detail.CVSS.Severity = getSeverityJa(c.Severity.String)
	}
	rf := []RelatedInfo{}
	dsR := db.GetRelated(id)
	if dsA != nil {
		for _, r := range dsR {
			rr := RelatedInfo{
				ID:        r.ID.String,
				Type:      r.Type.String,
				Name:      r.Name.String,
				VulinfoID: r.VulinfoID.String,
				Title:     r.Title.String,
				URL:       r.URL.String,
			}
			rf = append(rf, rr)
		}
	}
	detail.Relattions = rf

	if tf == nil {
		file, err := Assets.Open("/template-detail.md")
		if err != nil {
			return buf, err
		}
		defer file.Close()
		tf = file
	}
	tmpdata := &strings.Builder{}
	io.Copy(tmpdata, tf)
	t, err := template.New("Detail by Markdown").Parse(tmpdata.String())
	if err != nil {
		return buf, err
	}
	if err = t.Execute(buf, detail); err != nil {
		return nil, err
	}
	if f == FormHTML {
		return convHTML(buf), nil
	}
	return buf, nil
}

func getGetVulnviewByID(db *database.DB, id string) *database.Vulnview {
	v := db.GetVulnview(id)
	if v != nil {
		return v
	}

	if err := db.UpdateID(id); err != nil {
		db.GetLogger().Errorln(err)
		return nil
	}

	return db.GetVulnview(id)
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
