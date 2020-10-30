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
	file, err := os.Open(path)
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
	rules := getRuleFromFile("/Users/finger/code/mycode/nightingale/rss.json")
	resultCh := make(chan *ParseResult, 10)
	flagCh := make(chan int, len(rules))
	aliveCh := len(rules)
	for _, val := range rules {
		go func(r XmlRule) {
			ch := r.GenerateMsgChan()
			for {
				as, ok := <-ch
				if !ok {
					flagCh <- 1
					break
				}
				resultCh <- as
			}
		}(val)
	}

	for {
		select {
		case <-flagCh:
			aliveCh--
			if aliveCh <= 0 {
				flagCh = nil
				close(resultCh)
			}
		case val, ok := <-resultCh:
			if !ok {
				goto END
			}
			fmt.Println(val)
		}
	}
END:
	fmt.Println("all stop!!")
}

func TestCloseChan(t *testing.T) {
	fmt.Println("hh")
}
