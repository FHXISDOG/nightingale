package crawler

import (
	"fmt"
	"net/http"
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
		Url:             "https://coolshell.cn/feed",
	}
	return rule
}

func printRssContent(val *ParseResult) {
	fmt.Println("================")
	fmt.Println(val.Title)
	fmt.Println(val.Link)
	fmt.Println(val.Date)
	// fmt.Println(val.Content)
	fmt.Println("================")
}
func TestXmlParserChan(t *testing.T) {

}

func TestXmlParser(tt *testing.T) {
	rule := &XmlRule{
		ParentNode:      "//channel/item",
		TitleNode:       "//title",
		DescriptionNode: "//description",
		ContentNode:     "//content:encoded",
		LinkNode:        "//link",
		DateNode:        "//pubDate",
	}
	resp, err := http.Get("https://coolshell.cn/feed")
	if err != nil {
		fmt.Println("fuck")
	}
	defer resp.Body.Close()
	result := XmlAnalyser(resp.Body, rule)
	for _, val := range result {
		if val != nil {
			fmt.Println("================")
			fmt.Println(val.Title)
			fmt.Println(val.Link)
			fmt.Println(val.Date)
			// fmt.Println(val.Content)
			fmt.Println("================")
		}
	}
}
