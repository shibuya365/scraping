# GetYahoo!news

自動で Yahoo!ニュースのタイトルを取得する

# DEMO

ex.

## 国内

- [政府、追加経済対策に 73 兆円](https://news.yahoo.co.jp/pickup/6378687)
- [地方移住し家購入 政府支援へ](https://news.yahoo.co.jp/pickup/6378693)
- [医療深刻 自衛隊は最後の手段](https://news.yahoo.co.jp/pickup/6378688)
- [GoTo 利用者は感染疑い 2 倍 調査](https://news.yahoo.co.jp/pickup/6378685)

## ビジネス

- [中国系の太陽光 5 社 所得隠し](https://news.yahoo.co.jp/pickup/6378702)
- [GoTo 電子クーポン 不正続く](https://news.yahoo.co.jp/pickup/6378696)

## IT

- [流行語大賞どう話題に?分析](https://news.yahoo.co.jp/pickup/6378691)

# Features

新しいニュースのみ`今日の日付と時刻.md`にまとめる
そのニュースタイトルをクリックすると詳しいニュースを表示する
古いニュースは JSON ファイルの保存する

# Requirement

- go
- goquery

# Installation

go 言語をインストール後、

```bash
go get github.com/PuerkitoBio/goquery
```

# Usage

DEMO の実行方法など、"hoge"の基本的な使い方を説明する

```bash
git clone https://github.com/shibuya365/scraping.git
cd your_folder
go mod your_folder
go build
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

- shibuya365
- shibuya365days@gmail.com

# License

"GetYahoo!news" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
