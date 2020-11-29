package conf

import (
	"encoding/json"
	"os"
)

func WriteConfDB(news []Data) error {
	// ファイルの生成
	file, err := os.Create("conf/db2.json")
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
	// エンコードされたデータをコンソールに表示
	// fmt.Println(string(myJson))

	return nil
}
