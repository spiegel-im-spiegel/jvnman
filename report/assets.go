package report

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsddade591f965b4f2ded44aad3a09a7eb9aa490fa = "<table class=\"vulnview\">\n<thead>\n<tr>\n\t<th>ID</th>\n\t<th>タイトル</th>\n\t<th>概要</th>\n\t<th>想定される影響</th>\n\t<th>対策</th>\n\t<th>深刻度</th>\n\t<th>発見日</th>\n\t<th>公開日</th>\n\t<th>最終更新日</th>\n</tr>\n</thead>\n<tbody>\n{{ range . }}<tr>\n\t<td><a href=\"{{ .URI }}\">{{ .ID }}</a></td>\n\t<td>{{ .Title }}</td>\n\t<td>{{ .Description }}</td>\n\t<td>{{ .Impact }}</td>\n\t<td>{{ .Solution }}</td>\n\t<td>{{ .Severity }}</td>\n\t<td>{{ .DatePublic }}</td>\n\t<td>{{ .DatePublish }}</td>\n\t<td>{{ .DateUpdate }}</td>\n</tr>{{ end }}\n</tbody>\n</table>\n"
var _Assets1bd83af4ae07cdff7378f7b03586382dbb783224 = "| ID  | タイトル | 概要 | 想定される影響 | 対策 | 深刻度 | 発見日 | 公開日 | 最終更新日 |\n| --- | -------- | ---- | -------------- | ---- | ------ | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Description }} | {{ .Impact }} | {{ .Solution }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DatePublish }} | {{ .DateUpdate }} |\n{{ end }}\n"
var _Assetsdb17a49b5b675d815b56b4b2f5d324f1efc3fe5d = "<table class=\"vulnview\">\n<thead>\n<tr>\n\t<th>ID</th>\n\t<th>タイトル</th>\n\t<th>深刻度</th>\n\t<th>発見日</th>\n\t<th>最終更新日</th>\n</tr>\n</thead>\n<tbody>\n{{ range . }}<tr>\n\t<td><a href=\"{{ .URI }}\">{{ .ID }}</a></td>\n\t<td>{{ .Title }}</td>\n\t<td>{{ .Severity }}</td>\n\t<td>{{ .DatePublic }}</td>\n\t<td>{{ .DateUpdate }}</td>\n</tr>{{ end }}\n</tbody>\n</table>\n"
var _Assets6a710ffe05355fd8de925514584ec37062ef3c9c = "| ID  | タイトル | 深刻度 | 発見日 | 最終更新日 |\n| --- | -------- | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DateUpdate }} |\n{{ end }}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template-list-detail.html", "template-list-detail.md", "template-list.html", "template-list.md"}}, map[string]*assets.File{
	"/template-list.md": &assets.File{
		Path:     "/template-list.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525873729, 1525873729728152100),
		Data:     []byte(_Assets6a710ffe05355fd8de925514584ec37062ef3c9c),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1525001486, 1525001486915567000),
		Data:     nil,
	}, "/template-list-detail.html": &assets.File{
		Path:     "/template-list-detail.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525001902, 1525001902065543900),
		Data:     []byte(_Assetsddade591f965b4f2ded44aad3a09a7eb9aa490fa),
	}, "/template-list-detail.md": &assets.File{
		Path:     "/template-list-detail.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525002002, 1525002002100605800),
		Data:     []byte(_Assets1bd83af4ae07cdff7378f7b03586382dbb783224),
	}, "/template-list.html": &assets.File{
		Path:     "/template-list.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1525873664, 1525873664258700600),
		Data:     []byte(_Assetsdb17a49b5b675d815b56b4b2f5d324f1efc3fe5d),
	}}, "")
