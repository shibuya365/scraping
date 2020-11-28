package main

import (
	"fmt"
	"log"
	"net/http"

	"./conf" // 実装した設定ファイルパッケージの読み込み
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 設定ファイルを読み込む
	news, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	// ニュースの読み込み
	res, err := http.Get("https://news.yahoo.co.jp/")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// タイトルのみの抜き出し
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}
	section := doc.Find(".topicsList_main a")

	for i := 0; i < 8; i++ {
		line := section.Eq(i)
		text := line.Text()
		fmt.Println(i+1, text)

		// 新しいデータの作成
		var data conf.Data
		data.Title = text
		news = append(news, data)
		err = conf.WriteConfDB(news)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}
