package crawler

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/antchfx/xmlquery"
)

type Any interface{}

type Info struct {
	title, describle, content string
}

func TestXml(tt *testing.T) {
	resp, err := http.Get("https://coolshell.cn/feed")
	if err != nil {
		fmt.Println("fuck")
	}
	defer resp.Body.Close()

	result := getXmlElementByXmlQuery(resp.Body, "")
	fmt.Println(result)
}

func getXmlElementByXmlQuery(source io.Reader, rule string) []string {
	doc, err := xmlquery.Parse(source)
	result := make([]string, 20)
	if err != nil {
		fmt.Println("parse xml error", err)
	}
	list := xmlquery.Find(doc, "//channel/item")
	for _, val := range list {
		result = append(result, val.SelectElement("//title").InnerText())
	}
	return result
}

// get xml element by rule , source respone body ,rule -> rule
func getXmlElementByRule(source io.Reader, rule string) []string {
	result := make([]string, 20)
	decoder := xml.NewDecoder(source)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		fmt.Println(token)
		break
	}
	fmt.Println(decoder)
	return result
}