package conf

import (
	"os"
)

func WriteConfDB(news []string) error {
	f, err := os.Create("conf/news.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range news {
		f.WriteString(line + "\n")
	}
	return nil
}
