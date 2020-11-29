package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./conf" // 実装した設定ファイルパッケージの読み込み
	"github.com/PuerkitoBio/goquery"
)

var webs map[string]string = map[string]string{
	"TOP":      "https://news.yahoo.co.jp/",
	"DOMESTIC": "https://news.yahoo.co.jp/categories/domestic",
	"WORLD":    "https://news.yahoo.co.jp/categories/world",
	"Business": "https://news.yahoo.co.jp/categories/business",
	"IT":       "https://news.yahoo.co.jp/categories/it",
}

func main() {
	// 今日の日付
	day := time.Now()
	today := day.Format("2006-01-02")

	// 今日の新しいニュースのファイルの生成
	file, err := os.Create("news" + today + ".md")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	// 1行目を記入
	// _, err := file.WriteString("## news\n")
	file.WriteString("### " + today + " News\n")

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

	// ニュースの読み込み
	for _, v := range webs {
		fmt.Println(v)

	}

	res, err := http.Get("https://news.yahoo.co.jp/")
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
	section := doc.Find(".topicsList_main a")

	// 個別のニュースをチェック
	for i := 0; i < 8; i++ {
		line := section.Eq(i)

		// ニュースタイトル
		text := line.Text()

		// 属性を取得、existsはその属性が存在するか
		attr, _ := line.Attr("href")
		// fmt.Println(line.Attr("href"))

		// タイトルが既にあるか調べる
		_, ok := mapNews[text]

		// なかった場合の処理
		if !ok {
			// 今日のファイルへの書き込み
			// fmt.Fprint()の場合
			// _, err := fmt.Fprint(file, text)
			_, err := file.WriteString("- [" + text + "](" + attr + ")\n")
			if err != nil {
				log.Println(err)
			}

			// 新しいデータの作成
			var data conf.Data
			data.Date = today
			data.Title = text
			data.Attr = attr
			news = append(news, data)
		}
		// データの保存
		err = conf.WriteConfDB(news)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
