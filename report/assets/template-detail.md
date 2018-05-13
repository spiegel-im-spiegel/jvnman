# {{ .Info.Title }}

脆弱性対策情報ID: [{{ .Info.ID }}]({{ .Info.URI }})

{{ .Info.Description }}

## 想定される影響

{{ .Info.Impact }}

### 影響を受ける製品

{{ range .Affects }}- {{ .Name }} / {{ .ProductName }} {{ .VersionNumber }}
{{ end }}

### 深刻度

{{ with .CVSS }}{{ if .Severity }}{{ .Severity }}: {{ .BaseVector }}（{{ .BaseScore }}）{{ else }}CVSSv3 評価なし{{ end }}{{ end }}

## 対策

{{ .Info.Solution }}

## 関連情報

{{ range .Relattions }}- {{ if .Name }}{{ .Name }} {{ end }}[{{ .VulinfoID }}]({{ .URL }}) {{ if .Title }}{{ .Title }}{{ end }}
{{ end }}

## 更新情報

- 発見日 {{ .Info.DatePublic }}
- 公開日 {{ .Info.DatePublish }}
- 最終更新日 {{ .Info.DateUpdate }}

(Powerd by [JVN](https://jvn.jp/))
