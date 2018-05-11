package report

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets1bd83af4ae07cdff7378f7b03586382dbb783224 = "| ID  | タイトル | 概要 | 想定される影響 | 対策 | 深刻度 | 発見日 | 公開日 | 最終更新日 |\n| --- | -------- | ---- | -------------- | ---- | ------ | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Description }} | {{ .Impact }} | {{ .Solution }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DatePublish }} | {{ .DateUpdate }} |\n{{ end }}\n"
var _Assetsdb17a49b5b675d815b56b4b2f5d324f1efc3fe5d = "<table class=\"vulnview\">\n<thead>\n<tr>\n\t<th>ID</th>\n\t<th>タイトル</th>\n\t<th>深刻度</th>\n\t<th>発見日</th>\n\t<th>最終更新日</th>\n</tr>\n</thead>\n<tbody>\n{{ range . }}<tr>\n\t<td><a href=\"{{ .URI }}\">{{ .ID }}</a></td>\n\t<td>{{ .Title }}</td>\n\t<td>{{ .Severity }}</td>\n\t<td>{{ .DatePublic }}</td>\n\t<td>{{ .DateUpdate }}</td>\n</tr>{{ end }}\n</tbody>\n</table>\n"
var _Assets6a710ffe05355fd8de925514584ec37062ef3c9c = "| ID  | タイトル | 深刻度 | 発見日 | 最終更新日 |\n| --- | -------- | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DateUpdate }} |\n{{ end }}\n"
var _Assets2e246b9906b9934683a55e514f8552196d6256de = "# {{ .Info.Title }}\n\n脆弱性対策情報ID: [{{ .Info.ID }}]({{ .Info.URI }})\n\n{{ .Info.Description }}\n\n## 想定される影響\n\n{{ .Info.Impact }}\n\n### 影響を受ける製品\n\n{{ range .Affects }}- {{ .Name }} / {{ .ProductName }} {{ .VersionNumber }}\n{{ end }}\n\n### 深刻度\n\n{{ with .CVSS }}{{ if .Severity }}{{ .Severity }}: {{ .BaseVector }}（{{ .BaseScore }}）{{ else }}CVSSv3 評価なし{{ end }}{{ end }}\n\n## 対策\n\n{{ .Info.Solution }}\n\n## 関連情報\n\n{{ range .Relattions }}- {{ if .Name }}{{ .Name }} {{ end }}[{{ .VulinfoID }}]({{ .URL }}) {{ if .Title }}{{ .Title }}{{ end }}\n{{ end }}\n\n## 更新情報\n\n- 発見日 {{ .Info.DatePublic }}\n- 公開日 {{ .Info.DatePublish }}\n- 最終更新日 {{ .Info.DateUpdate }}\n"
var _Assetsddade591f965b4f2ded44aad3a09a7eb9aa490fa = "<table class=\"vulnview\">\n<thead>\n<tr>\n\t<th>ID</th>\n\t<th>タイトル</th>\n\t<th>概要</th>\n\t<th>想定される影響</th>\n\t<th>対策</th>\n\t<th>深刻度</th>\n\t<th>発見日</th>\n\t<th>公開日</th>\n\t<th>最終更新日</th>\n</tr>\n</thead>\n<tbody>\n{{ range . }}<tr>\n\t<td><a href=\"{{ .URI }}\">{{ .ID }}</a></td>\n\t<td>{{ .Title }}</td>\n\t<td>{{ .Description }}</td>\n\t<td>{{ .Impact }}</td>\n\t<td>{{ .Solution }}</td>\n\t<td>{{ .Severity }}</td>\n\t<td>{{ .DatePublic }}</td>\n\t<td>{{ .DatePublish }}</td>\n\t<td>{{ .DateUpdate }}</td>\n</tr>{{ end }}\n</tbody>\n</table>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template-detail.md", "template-list-detail.html", "template-list-detail.md", "template-list.html", "template-list.md"}}, map[string]*assets.File{
	"/template-list-detail.html": &assets.File{
		Path:     "/template-list-detail.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525133899, 1525133899254034300),
		Data:     []byte(_Assetsddade591f965b4f2ded44aad3a09a7eb9aa490fa),
	}, "/template-list-detail.md": &assets.File{
		Path:     "/template-list-detail.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525133899, 1525133899254034300),
		Data:     []byte(_Assets1bd83af4ae07cdff7378f7b03586382dbb783224),
	}, "/template-list.html": &assets.File{
		Path:     "/template-list.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525910718, 1525910718886121200),
		Data:     []byte(_Assetsdb17a49b5b675d815b56b4b2f5d324f1efc3fe5d),
	}, "/template-list.md": &assets.File{
		Path:     "/template-list.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525910718, 1525910718901745300),
		Data:     []byte(_Assets6a710ffe05355fd8de925514584ec37062ef3c9c),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1525934422, 1525934422033163000),
		Data:     nil,
	}, "/template-detail.md": &assets.File{
		Path:     "/template-detail.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1526027415, 1526027415741484200),
		Data:     []byte(_Assets2e246b9906b9934683a55e514f8552196d6256de),
	}}, "")
