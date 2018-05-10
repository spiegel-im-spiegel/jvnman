package report

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/spiegel-im-spiegel/jvnman/database"
	"gopkg.in/russross/blackfriday.v2"
)

type VulnDetail struct {
	Info VulnInfo
}

func Detail(db *database.DB, id string, f Format) (io.Reader, error) {
	buf := &bytes.Buffer{}
	detail := VulnDetail{}
	v := db.GetVulnview(id)
	if v == nil {
		return buf, nil
	}
	detail.Info.ID = v.ID.String
	detail.Info.Title = v.Title.String
	detail.Info.Description = v.Description.String
	detail.Info.URI = v.URI.String
	detail.Info.Impact = v.Impact.String
	detail.Info.Solution = v.Solution.String
	detail.Info.Severity = fmt.Sprintf("%v (%.1f)", getSeverityJa(v.CVSSSeverity.String), v.CVSSScore.Float64)
	detail.Info.DatePublic = v.GetDatePublic().Format("2006年1月2日")
	detail.Info.DatePublish = v.GetDatePublish().Format("2006年1月2日")
	detail.Info.DateUpdate = v.GetDateUpdate().Format("2006年1月2日")

	tf, err := Assets.Open("/template-detail.md")
	if err != nil {
		return buf, err
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

func convHTML(md *bytes.Buffer) io.Reader {
	//HTMLFlags and Renderer
	htmlFlags := blackfriday.CommonHTMLFlags         //UseXHTML | Smartypants | SmartypantsFractions | SmartypantsDashes | SmartypantsLatexDashes
	htmlFlags |= blackfriday.FootnoteReturnLinks     //Generate a link at the end of a footnote to return to the source
	htmlFlags |= blackfriday.SmartypantsAngledQuotes //Enable angled double quotes (with Smartypants) for double quotes rendering
	htmlFlags |= blackfriday.SmartypantsQuotesNBSP   //Enable French guillemets 損 (with Smartypants)
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: htmlFlags, Title: "", CSS: ""})

	//Extensions
	extFlags := blackfriday.CommonExtensions //NoIntraEmphasis | Tables | FencedCode | Autolink | Strikethrough | SpaceHeadings | HeadingIDs | BackslashLineBreak | DefinitionLists
	extFlags |= blackfriday.Footnotes        //Pandoc-style footnotes
	extFlags |= blackfriday.HeadingIDs       //specify heading IDs  with {#id}
	extFlags |= blackfriday.Titleblock       //Titleblock ala pandoc
	extFlags |= blackfriday.DefinitionLists  //Render definition lists

	html := blackfriday.Run(md.Bytes(), blackfriday.WithExtensions(extFlags), blackfriday.WithRenderer(renderer))
	return bytes.NewReader(html)
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
