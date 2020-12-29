package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
	var wg sync.WaitGroup
	// カテゴリごとに繰り返す
	for cate, web := range webs {
		wg.Add(1)
		go getCate(cate, web, &wg)
	}
	// 並列処理が終わるのをまつ
	wg.Wait()
}

func getCate(cate string, web string, wg *sync.WaitGroup) {
	defer wg.Done()
	// 設定ファイルを読み込む
	news, err := conf.ReadConfDB(cate)
	// もしファイルがなかったら
	if err != nil {
		fmt.Println(err.Error())
	}

	// カテゴリのジャンルを表示へ
	ans := "# " + cate + "\n"

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

	section := doc.Find(".sc-DNdyV a")

	// 個別のニュースをチェック
	section.Each(func(i int, line *goquery.Selection) {

		// ニュースタイトルを取得
		title := line.Text()

		// href属性を取得
		attr, _ := line.Attr("href")

		// ニュースが既出が調べる
		var app bool
		for _, new := range news {
			if title == "もっと見る" || title == new {
				app = true
				break
			}
		}
		// もしなかったら
		if !app {
			// コンソールへ新しいニュースのタイトルのみ出力
			ans += title + "		:" + attr + "\n"
			// 新しいニュースタイトルを追加
			news = append(news, title)
		}
	})
	// ニュースを表示
	fmt.Println(ans)

	// データの保存
	if len(news) > 12 {
		news = news[len(news)-19:]
	}
	err = conf.WriteConfDB(cate, news)
	if err != nil {
		fmt.Println(err.Error())
	}
}
