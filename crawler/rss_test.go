package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Any interface{}

type Info struct {
	title, describle, content string
}

func generateRule() *XmlRule {
	rule := &XmlRule{
		ParentNode:      "//channel/item",
		TitleNode:       "//title",
		DescriptionNode: "//description",
		ContentNode:     "//content:encoded",
		LinkNode:        "//link",
		DateNode:        "//pubDate",
		Rule: &Rule{
			Url:           "https://coolshell.cn/feed",
			CanPage:       false,
			RequestMethod: GET,
		},
	}
	return rule
}

func getRuleFromFile(path string) []XmlRule {
	file, err := os.Open("/Users/finger/code/mycode/nightingale/rss.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	res := make([]XmlRule, 0)
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &res)
	return res
}

func TestXmlParserChan(t *testing.T) {
	rules := getRuleFromFile("/Users/finger/code")
	resultCh := make(chan *ParseResult, 2)
	for _, val := range rules {
		go val.GenerateResult(resultCh)
	}
	for val := range resultCh {
		fmt.Println(val)
	}
}

// test json file to struct
// func TestReadJson(t *testing.T) {
// 	file, err := os.Open("/Users/finger/code/mycode/nightingale/rss.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()
// 	res := make([]XmlRule, 0)
// 	byteValue, _ := ioutil.ReadAll(file)
// 	fmt.Println(string(byteValue))
// 	json.Unmarshal(byteValue, &res)
// 	fmt.Println(res)
// }
