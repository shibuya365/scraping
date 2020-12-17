package conf

import (
	"encoding/json"
	"os"
)

func WriteConfDB(news []string) error {
	// ファイルの生成
	file, err := os.Create("conf/db.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// JSONファイルへ変換
	myJson, err := json.Marshal(&news)
	if err != nil {
		return err
	}

	// 正常に生成されたファイルに書き込み
	file.Write(([]byte)(string(myJson)))

	return nil
}
