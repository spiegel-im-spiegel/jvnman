module github.com/spiegel-im-spiegel/jvnman

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jessevdk/go-assets v0.0.0-20160921144138-4f4301a06e15
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20180821113735-8b31f9c59b0f // indirect
	github.com/mattn/go-sqlite3 v1.9.0
	github.com/mitchellh/go-homedir v1.0.0
	github.com/pkg/errors v0.8.0
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.2.1
	github.com/spiegel-im-spiegel/go-cvss v0.1.1
	github.com/spiegel-im-spiegel/go-myjvn v0.4.0
	github.com/spiegel-im-spiegel/gocli v0.8.0
	github.com/spiegel-im-spiegel/logf v0.2.3
	golang.org/x/text v0.3.0
	gopkg.in/Masterminds/squirrel.v1 v1.0.0-20170825200431-a6b93000bd21
	gopkg.in/gorp.v2 v2.0.0
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1
