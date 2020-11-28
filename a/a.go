package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Blog struct {
	Title string `json:"title"`
}

func main() {
	// JSONの読み込み
	jsonFromFile, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get("https://news.yahoo.co.jp/")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	section := doc.Find(".topicsList_main a")

	// section = section.Eq(1)
	// text := section.Text()
	// fmt.Println(text)
	for i := 0; i < 8; i++ {
		line := section.Eq(i)
		text := line.Text()
		fmt.Println(i+1, text)
	}

	// for i, v := range section {
	// text := section.Text()
	// fmt.Println(i+1, text)
	// }
}
