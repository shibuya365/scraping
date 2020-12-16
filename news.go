package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/get_yahoo_news/conf"
	"github.com/PuerkitoBio/goquery"
)

// Yahoo!のニュースサイトのアドレス
var webs map[string]string = map[string]string{
	"国内":   "https://news.yahoo.co.jp/categories/domestic",
	"国際":   "https://news.yahoo.co.jp/categories/world",
	"ビジネス": "https://news.yahoo.co.jp/categories/business",
	"IT":   "https://news.yahoo.co.jp/categories/it",
}

func main() {
	// 今日の日付
	day := time.Now()
	today := day.Format("2006-01-02")
	time := day.Format("2006-01-02 03:04:05")

	// 今日の新しいニュースのファイルの生成
	file, err := os.Create("news" + time + ".md")
	// file, err := os.Create("news.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	// 1行目にタイトルを記入
	file.WriteString("# " + today + " News\n")

	// 設定ファイルを読み込む
	news, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 古いニュースをmapへ
	mapNews := make(map[string]bool)
	for i, _ := range news {
		mapNews[news[i].Title] = true
	}

	// カテゴリごとに繰り返す
	for i, web := range webs {
		// カテゴリをMDへ
		file.WriteString("## " + i + "\n")

		// ニュースの読み込み
		res, err := http.Get(web)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
		}

		// タイトルの部分の抜き出し
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Println(err)
		}
		section := doc.Find(".sc-cgzHhG a")

		// 個別のニュースをチェック
		section.Each(func(i int, line *goquery.Selection) {

			// ニュースタイトルを取得
			text := line.Text()

			// href属性を取得
			attr, _ := line.Attr("href")

			// タイトルが既にあるか調べる
			_, ok := mapNews[text]

			// なかった場合の処理
			if !ok {
				// コンソールへ新しいニュースのタイトルのみ出力
				fmt.Println(text)
				// 今日のファイルへの書き込み
				_, err := file.WriteString("- [" + text + "](" + attr + ")\n")
				if err != nil {
					log.Println(err)
				}

				// 新しいデータの作成
				var data conf.Data
				data.Title = text
				news = append(news, data)
			}
		})
	}

	// 全てのニュース追加後データの保存
	err = conf.WriteConfDB(news)
	if err != nil {
		fmt.Println(err.Error())
	}

}
