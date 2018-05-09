| ID  | タイトル | 深刻度 | 発見日 | 最終更新日 |
| --- | -------- | ------ | ------ | ---------- |
{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DateUpdate }} |
{{ end }}
