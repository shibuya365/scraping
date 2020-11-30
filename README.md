# GetYahoo!news
 
自動でYahoo!ニュースのタイトルを取得する
# DEMO
`golang`をインストールした状態で
```bash
go get github.com/PuerkitoBio/goquery
git clone https://github.com/shibuya365/scraping.git
cd examples
go run .
```
# Features
既に一度見たファイルをJSONファイルの保存し新しいニュースのみ`今日の日付と時刻.md`にまとめる
そのニュースタイトルをクリックすると詳しいニュースを表示する

# Requirement
 
* go
* goquery
 
# Installation
 go言語をインストール後、
```bash
go get github.com/PuerkitoBio/goquery
```
 
# Usage
 
DEMOの実行方法など、"hoge"の基本的な使い方を説明する
 
```bash
git clone https://github.com/shibuya365/scraping.git
cd examples
go run .
```
 
# Note
 
`main.go`の以下の部分にお好きなカテゴリーを登録する
```go
// Yahoo!のニュースサイトのアドレス
var webs map[string]string = map[string]string{
	"国内":   "https://news.yahoo.co.jp/categories/domestic",
	"国際":   "https://news.yahoo.co.jp/categories/world",
	"ビジネス": "https://news.yahoo.co.jp/categories/business",
	"IT":   "https://news.yahoo.co.jp/categories/it",
}
```

# Author
 
* shibuya365
* shibuya365days@gmail.com
 
# License
 
"GetYahoo!news" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).