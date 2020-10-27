package crawler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/antchfx/xmlquery"
)

// resulr of after analyse
type ParseResult struct {
	Title, Author, Description, Content, Link, Source string
	Date                                              string
}

// rule of rss source
type XmlRule struct {
	ParentNode, TitleNode, DescriptionNode, ContentNode, LinkNode, DateNode string
	Url                                                                     string
}

// convert rss content to my result
func (rule *XmlRule) GetResult(val *xmlquery.Node) *ParseResult {
	title := val.SelectElement(rule.TitleNode).InnerText()
	description := val.SelectElement(rule.DescriptionNode).InnerText()
	content := val.SelectElement(rule.ContentNode).InnerText()
	link := val.SelectElement(rule.LinkNode).InnerText()
	date := val.SelectElement(rule.DateNode).InnerText()

	return &ParseResult{
		Title:       title,
		Content:     content,
		Description: description,
		Date:        date,
		Link:        link,
	}
}

// http get rss content
func (rule *XmlRule) RequestRss() io.ReadCloser {
	url := rule.Url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http error !!!")
	}
	return resp.Body
}

func (rule *XmlRule) GenerateResult(outChan chan *ParseResult) {
	source := rule.RequestRss()
	defer source.Close()
	doc, _ := xmlquery.Parse(source)
	parentNode := xmlquery.Find(doc, rule.ParentNode)

	contentLen := len(parentNode)

	for _, val := range parentNode {
		go func(v *xmlquery.Node) {
			result := rule.GetResult(v)
			outChan <- result
			finishLenCh <- 1
		}(val)
	}
}
