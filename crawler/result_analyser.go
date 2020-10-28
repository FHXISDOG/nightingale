package crawler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/antchfx/xmlquery"
)

const (
	GET  = 1
	POST = 2
)

const (
	URL  = 1
	BODY = 2
)

// resulr of after analyse
type ParseResult struct {
	Title, Author, Description, Content, Link, Source string
	Date                                              string
}

type Rule struct {
	Url           string
	RequestMethod int
	CanPage       bool
	InitPage      int
}

// rule of rss source
type XmlRule struct {
	*Rule                                                                   `json:"HttpRule"`
	ParentNode, TitleNode, DescriptionNode, ContentNode, LinkNode, DateNode string
	Body                                                                    string
}

func (val ParseResult) String() string {
	res := "===========================\n"
	res += val.Title + "\n"
	res += val.Link + "\n"
	res += val.Date + "\n"
	res += "===========================\n"
	return res
}

func (rule *Rule) generateHttpInitmsg() *HttpInitMsg {
	return &HttpInitMsg{
		Url: rule.Url,
	}
}

func (rule *Rule) generateNextpage(initMsg *HttpInitMsg, page int) {
	initMsg.Url = initMsg.Url
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
func (rule *XmlRule) RequestRss(initMsg *HttpInitMsg) io.ReadCloser {
	url := initMsg.Url
	var resp *http.Response
	var err error
	switch rule.RequestMethod {
	case GET:
		resp, err = http.Get(url)
	case POST:
		fmt.Println("UNKNOWN REQUEST METHOD")
	default:
		fmt.Println("UNKNOWN REQUEST METHOD")
	}
	if err != nil {
		fmt.Println("http error !!!")
	}
	return resp.Body
}

func (rule *XmlRule) GenerateResult(outChan chan *ParseResult) {
	initMsg := rule.generateHttpInitmsg()
	canPage := rule.CanPage
	var page int
	if canPage {
		page = rule.InitPage
	}
	for {
		source := rule.RequestRss(initMsg)
		defer source.Close()
		doc, _ := xmlquery.Parse(source)
		parentNode := xmlquery.Find(doc, rule.ParentNode)

		for _, val := range parentNode {
			result := rule.GetResult(val)
			outChan <- result
		}

		if !canPage {
			break
		}
		page++
		rule.generateNextpage(initMsg, page)
	}
}
