# [jvnman] - JVN Vulnerability Data Management

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/jvnman.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/jvnman)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/jvnman/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/jvnman.svg)](https://github.com/spiegel-im-spiegel/jvnman/releases/latest)

[jvnman] は [JVN] が提供する「[脆弱性対策情報共有フレームワーク]」を使った脆弱性情報の収集・管理ツールです。

## [jvnman] のインストール

[jvnman] のインストールは `go get` コマンドで自動的にビルド＆インストールしてくれます。

```
$ go get -u github.com/spiegel-im-spiegel/jvnman
```

なお，内部で [github.com/mattn/go-sqlite3] パッケージを使用しているため，ビルド時に [GCC] C コンパイラが必要です。
Windows 環境の場合は別途 [GCC] のインストールが必要です。
Windows 環境で [GCC] のインストールを行う場合は [MinGW-w64] を利用することを強くおすすめします。

- [MinGW-w64 を導入する — しっぽのさきっちょ | text.Baldanders.info](http://text.baldanders.info/remark/2018/03/mingw-w64/)
- [Go 言語で SQLite を使う（Windows 向けの紹介） — プログラミング言語 Go | text.Baldanders.info](http://text.baldanders.info/golang/sqlite-with-golang-in-windows/)

なお `go get` コマンドでビルドに失敗する場合は [dep] コマンドで vendoring すると上手くいくかもしれません。

```
$ dep ensure
```

## [jvnman] の使い方

```
$ jvnman -h
JVN database management

Usage:
  jvnman [flags]
  jvnman [command]

Available Commands:
  help        Help about any command
  info        Output vulnerability information
  init        Initialize vulnerability database
  list        List vulnerability data
  update      Update vulnerability database
  version     Print the version number

Flags:
      --config string     config file (default $HOME/.jvnman.yaml)
      --dbfile string     database file path (default "./jvndb.sqlite3")
  -h, --help              help for jvnman
      --logfile string    logfile path (default standard error)
      --loglevel string   log level: trace/debug/info/warn/error/fatal (default "error")

Use "jvnman [command] --help" for more information about a command.
```

### データベースの初期化

[jvnman] はローカルに [SQLite] データベース・ファイルを作成し情報の蓄積・管理を行います。
まず最初に `jvnman init` コマンドでデータベースの初期化を行います。

`jvnman init` コマンドの usage は以下の通り。

```
$ jvnman init -h
Initialize vulnerability database

Usage:
  jvnman init [flags]

Flags:
  -h, --help   help for init

Global Flags:
      --config string     config file (default $HOME/.jvnman.yaml)
      --dbfile string     database file path (default "./jvndb.sqlite3")
      --logfile string    logfile path (default standard error)
      --loglevel string   log level: trace/debug/info/warn/error/fatal (default "error")
```

初期化を行う場合は引数なしで `jvnman init` コマンドを実行します。

```
$ jvnman init
```

これでカレントディレクトリに `jvndb.sqlite3` ファイルが作成されます。
また `--dbfile` オプションでパスを直接指定することもできます。

 [SQLite] データベース・ファイルを初期化して作り直したい場合も `jvnman init` コマンドで初期化できます。

データベースのテーブル仕様はリポジトリの [jvnman/docs/](https://github.com/spiegel-im-spiegel/jvnman/tree/master/docs) フォルダ以下にあります。

### 脆弱性情報の蓄積・更新

[JVN] のデータベースから脆弱性情報を取得するには `jvnman update` コマンドを用います。

```
$ jvnman update -h
Update vulnerability database

Usage:
  jvnman update [flags]

Flags:
  -h, --help             help for update
  -k, --keyword string   keyword for filtering
  -m, --month            get the data for the past month

Global Flags:
      --config string     config file (default $HOME/.jvnman.yaml)
      --dbfile string     database file path (default "./jvndb.sqlite3")
      --logfile string    logfile path (default standard error)
      --loglevel string   log level: trace/debug/info/warn/error/fatal (default "error")
```

引数なしで

```
$ jvnman update
```

とすると1週間以内に公開・更新された脆弱性情報を取得します。
既に蓄積された除法がある場合は前回の最新更新日以降に公開・更新された脆弱性情報を取得します。
`-m` オプションを指定すると1ヶ月以内に公開・更新された脆弱性情報を取得します。

`-k` オプションでキーワードを指定するとキーワードに合致する情報のみを取得します。
たとえば， Java 関連の情報のみを収集したければ

```
$ jvnman update -k java
```

として情報を絞り込むことができます。
なお，大文字小文字の区別はしません。

### 一覧表の出力

`jvnman list` コマンドを用いて蓄積した脆弱性情報を一覧表示することができます。

```
$ jvnman list -h
List vulnerability data

Usage:
  jvnman list [flags]

Flags:
  -c, --cve string        CVE-ID (see https://cve.mitre.org/) for filtering
  -f, --form string       output format: html/markdown/csv (default "markdown")
  -h, --help              help for list
  -p, --product string    product name for filtering
  -r, --range int         list the data for the past days (default 3)
  -s, --score float       minimum score of CVSS for filtering
  -t, --template string   template file path

Global Flags:
      --config string     config file (default $HOME/.jvnman.yaml)
      --dbfile string     database file path (default "./jvndb.sqlite3")
      --logfile string    logfile path (default standard error)
      --loglevel string   log level: trace/debug/info/warn/error/fatal (default "error")
```

出力フォーマットには html, markdown, csv を指定できます。
既定は markdown です。
一覧表は標準出力に出力されます。
出力時の文字エンコーディングは UTF-8 です。

```
$ jvnman list -f csv > list.csv
```

コマンドラインでいくつかのフィルタリング条件を指定できます。
オプションと条件の対応は以下のとおりです。

| オプション | 既定値            | 内容                                             |
| ---------- | ----------------- | ------------------------------------------------ |
| `-c`       | 指定なし          | CVE-ID でフィルタリングします                    |
| `-p`       | 指定なし          | ベンダ名または製品名でフィルタリングします       |
| `-r`       | 現時点から3日以内 | 出力範囲を最終更新日をキーにフィルタリングします |
| `-s`       | 0 (指定なし)      | CVSS 基本評価値の下限値でフィルタリングします    |

各条件は論理積（AND）で効きます。
たとえば

```
$ jvnman list -p java -s 7.0 -r 30
```

とすれば，過去30日以内に更新された情報のうち Java 関連でかつ CVSS 基本評価値が 7.0 以上のものに絞って出力されます。

また，あらかじめ CVE-ID の値がわかっているのであれば

```
$ $ jvnman list -c CVE-2018-2783 -r 30
```

という感じで絞り込むことができます。

`-t` オプションで一覧を整形する際のテンプレートファイルを指定できます。

```
$ jvnman list -t template-list.md
```

[jvnman] が使用している標準のテンプレートファイルはリポジトリの [jvnman/report/assets/](https://github.com/spiegel-im-spiegel/jvnman/tree/master/report/assets) フォルダにあります。
参考にしてください。

### 脆弱性情報の出力

`jvnman info` コマンドを用いて指定した脆弱性情報 ID の脆弱性情報を帳票として出力することができます。

```
$ jvnman info -h
Output vulnerability information

Usage:
  jvnman info [flags] <JVN Vulnerability ID>

Flags:
  -f, --form string       output format: html/markdown (default "markdown")
  -h, --help              help for info
  -t, --template string   template file path

Global Flags:
      --config string     config file (default $HOME/.jvnman.yaml)
      --dbfile string     database file path (default "./jvndb.sqlite3")
      --logfile string    logfile path (default standard error)
      --loglevel string   log level: trace/debug/info/warn/error/fatal (default "error")
```

出力フォーマットには html, markdown を指定できます。
既定は markdown です。
一覧表は標準出力に出力されます。
出力時の文字エンコーディングは UTF-8 です。

```
$ jvnman info -f html JVNDB-2018-002862  > JVNDB-2018-002862.html
```

ローカルの [SQLite] データベース・ファイルに情報がない場合は [JVN] のデータベースにアクセスして情報を取得します。

`-t` オプションで情報を整形する際のテンプレートファイルを指定できます。

```
$ jvnman info -t template-info.md JVNDB-2018-002862
```

[jvnman] が使用している標準のテンプレートファイルはリポジトリの [jvnman/report/assets/](https://github.com/spiegel-im-spiegel/jvnman/tree/master/report/assets) フォルダにあります。
参考にしてください。

### ログ出力について

[jvnman] は実行時の状況をログ出力します。
ログの出力先ファイルと出力レベルは `--logfile` および `--loglevel` オプションで設定できます。

出力先の既定は標準エラー出力です。
ログファイルのハンドリングについては [github.com/lestrrat-go/file-rotatelogs] パッケージを用いているため `jvnman-%Y%m%d%H%M.log` のような指定も可能です。

出力レベルは trace, debug, info, warn, error, fatal の6段階で指定します。
既定値は error です。
trace を指定するとデバッグ用に SQL 文を出力するためログの量が大量になります。
ご注意ください。

```
$ jvnman update --logfile jvnman-%Y%m%d%H%M.log --loglevel trace
```

### 環境設定ファイル

[jvnman] の起動時，ホーム・ディレクトリに `.jvnman.yaml` ファイルがあれば環境設定ファイルとして読み込みます。
環境設定ファイルには `dbfile`, `logfile`, `loglevel` 各オプションを YAML 形式で記述します。

```yaml
dbfile: /path/to/jvndb.sqlite3
logfile: /path/to/jvnman-%Y%m%d%H%M.log
loglevel: trace
```

[jvnman] 起動時の `--config` オプションで環境設定ファイルのパスを指定することもできます。
環境設定ファイルの記述よりコマンドラインの指定のほうが優先されます。

## フィードバック

フィードバックはお気軽にどうぞ。
日本語で大丈夫です。

## 参考情報

- [MyJVN API に関する覚え書き — しっぽのさきっちょ | text.Baldanders.info](http://text.baldanders.info/remark/2018/03/myjvn-api/)
- [spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang](https://github.com/spiegel-im-spiegel/go-myjvn)

[jvnman]: https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Data management"
[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[github.com/mattn/go-sqlite3]: https://github.com/mattn/go-sqlite3 "mattn/go-sqlite3: sqlite3 driver for go using database/sql"
[github.com/lestrrat-go/file-rotatelogs]: https://github.com/lestrrat-go/file-rotatelogs "lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go"
[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[MinGW-w64]: http://mingw-w64.org/ "Mingw-w64 - GCC for Windows 64 & 32 bits [mingw-w64]"
[dep]: https://golang.github.io/dep/ "dep · Dependency management for Go"
[SQLite]: https://www.sqlite.org/
