package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mycode/nightingale/crawler"
	"os"
	"testing"
)

func TestReadFromJson(t *testing.T) {
	file, err := os.Open("/Users/finger/code/mycode/nightingale/rss.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	res := make([]crawler.XmlRule, 0)
	byteValue, _ := ioutil.ReadAll(file)
	fmt.Println(string(byteValue))
	json.Unmarshal(byteValue, &res)
	fmt.Println(res)
}
