package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mycode/nightingale/crawler"
	"os"
)

func ReadRss(path string) []crawler.XmlRule {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	res := make([]crawler.XmlRule, 0)
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &res)
	return res
}
