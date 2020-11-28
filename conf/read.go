package conf // 独自の設定ファイルパッケージ

import (
	"encoding/json" // DB設定の構造体
	"io/ioutil"
)

type Data struct {
	Title string `json:"title"`
}

var news []Data

// DB設定読み込み関数
func ReadConfDB() ([]Data, error) {

	// 設定ファイルを読み込む
	myJson, err := ioutil.ReadFile("conf/db.json")
	if err != nil {
		return news, err
	}

	err = json.Unmarshal(myJson, &news)

	if err != nil {
		return news, err
	}

	// fmt.Printf("News : %v", news)
	return news, nil
}
