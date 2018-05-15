package report

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets6a710ffe05355fd8de925514584ec37062ef3c9c = "| ID  | タイトル | 概要 | 想定される影響 | 対策 | CVSSv3ベクタ | 深刻度 | 発見日 | 公開日 | 最終更新日 |\n| --- | -------- | ---- | -------------- | ---- | ------------- | ------ | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Description }} | {{ .Impact }} | {{ .Solution }} | {{ .CVSSVector }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DatePublish }} | {{ .DateUpdate }} |\n{{ end }}\n\n(Powerd by [JVN](https://jvn.jp/))\n"
var _Assetse7f56da64dd0492253908aa8f211da1ca94f89dd = "| {{ .BaseMetrics }} | {{ .MetricValue }} |\n|--------|-------|\n| {{ .AVName }} | {{ .AVValue }} |\n| {{ .ACName }} | {{ .ACValue }} |\n| {{ .PRName }} | {{ .PRValue }} |\n| {{ .UIName }} | {{ .UIValue }} |\n| {{ .SName }} | {{ .SValue }} |\n| {{ .CName }} | {{ .CValue }} |\n| {{ .IName }} | {{ .IValue }} |\n| {{ .AName }} | {{ .AValue }} |\n"
var _Assets2e246b9906b9934683a55e514f8552196d6256de = "# {{ .Info.Title }}\n\n脆弱性対策情報ID: [{{ .Info.ID }}]({{ .Info.URI }})\n\n{{ .Info.Description }}\n\n## 想定される影響\n\n{{ .Info.Impact }}\n\n### 影響を受ける製品\n\n{{ range .Affects }}- {{ .Name }} / {{ .ProductName }} {{ .VersionNumber }}\n{{ end }}\n\n### 深刻度\n\n{{ with .CVSS }}{{ if .Severity }}{{ .Severity }}: {{ .BaseVector }}（{{ .BaseScore }}）\n\n{{ .BaseReport }}{{ else }}CVSSv3 評価なし{{ end }}{{ end }}\n\n## 対策\n\n{{ .Info.Solution }}\n\n## 関連情報\n\n{{ range .Relattions }}- {{ if .Name }}{{ .Name }} {{ end }}[{{ .VulinfoID }}]({{ .URL }}) {{ if .Title }}{{ .Title }}{{ end }}\n{{ end }}\n\n## 更新情報\n\n- 発見日 {{ .Info.DatePublic }}\n- 公開日 {{ .Info.DatePublish }}\n- 最終更新日 {{ .Info.DateUpdate }}\n\n(Powerd by [JVN](https://jvn.jp/))\n"
var _Assets12ec10cd20a0fc6f3fd5d9740c76f5d14a1e1bbd = "| ID  | タイトル | 深刻度 | 発見日 | 最終更新日 |\n| --- | -------- | ------ | ------ | ---------- |\n{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DateUpdate }} |\n{{ end }}\n\n(Powerd by [JVN](https://jvn.jp/))\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"cvss.md", "template-detail.md", "template-list-simple.md", "template-list.md"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1526359405, 1526359405715094400),
		Data:     nil,
	}, "/cvss.md": &assets.File{
		Path:     "/cvss.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1526368907, 1526368907868196100),
		Data:     []byte(_Assetse7f56da64dd0492253908aa8f211da1ca94f89dd),
	}, "/template-detail.md": &assets.File{
		Path:     "/template-detail.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1526361487, 1526361487558370900),
		Data:     []byte(_Assets2e246b9906b9934683a55e514f8552196d6256de),
	}, "/template-list-simple.md": &assets.File{
		Path:     "/template-list-simple.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1526257608, 1526257608136299100),
		Data:     []byte(_Assets12ec10cd20a0fc6f3fd5d9740c76f5d14a1e1bbd),
	}, "/template-list.md": &assets.File{
		Path:     "/template-list.md",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1526257608, 1526257608136299100),
		Data:     []byte(_Assets6a710ffe05355fd8de925514584ec37062ef3c9c),
	}}, "")
