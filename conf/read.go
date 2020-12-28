package conf // 独自の設定ファイルパッケージ

import (
	"bufio"
	"os"
)

var news []string

// DB設定読み込み関数
func ReadConfDB() ([]string, error) {
	// 設定ファイルを読み込む
	f, err := os.Open("conf/news.txt")
	if err != nil {
		return news, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// ここで一行ずつ処理
		news = append(news, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return news, err
	}

	return news, nil
}
