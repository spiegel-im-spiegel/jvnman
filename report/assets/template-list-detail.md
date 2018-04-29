| ID  | タイトル | 概要 | 想定される影響 | 対策 | 深刻度 | 発見日 | 公開日 | 最終更新日 |
| --- | -------- | ---- | -------------- | ---- | ------ | ------ | ------ | ---------- |
{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Description }} | {{ .Impact }} | {{ .Solution }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DatePublish }} | {{ .DateUpdate }} |
{{ end }}
