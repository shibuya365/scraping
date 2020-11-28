package main

import (
	"fmt"

	"./conf" // 実装した設定ファイルパッケージの読み込み
)

func main() {

	// 設定ファイルを読み込む
	news, err := conf.ReadConfDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	var data conf.Data
	data.Title = "あいうえお"
	news = append(news, data)
	err = conf.WriteConfDB(news)
	if err != nil {
		fmt.Println(err.Error())
	}
}
