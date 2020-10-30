package crawler

import (
	"fmt"
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
	res += val.Description + "\n"
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
func (rule *XmlRule) GeneratePageMsg(page int, httpMsg *HttpInitMsg) {

}

// http get rss content
func (rule *XmlRule) RequestRss(initMsg *HttpInitMsg) []*xmlquery.Node {
	url := initMsg.Url
	var resp *http.Response
	var err error
	switch rule.RequestMethod {
	case GET:
		resp, err = http.Get(url)
	case POST:
		//todo handle post request
		fmt.Println("UNKNOWN REQUEST METHOD")
	default:
		fmt.Println("UNKNOWN REQUEST METHOD")
	}
	if err != nil {
		fmt.Println("http error !!!")
	}
	source := resp.Body
	defer source.Close()
	doc, _ := xmlquery.Parse(source)
	parentNode := xmlquery.Find(doc, rule.ParentNode)
	return parentNode
}

// loop generate http request msg
func (rule *XmlRule) GenerateParentchan() chan []*xmlquery.Node {
	ch := make(chan []*xmlquery.Node)
	initMsg := rule.generateHttpInitmsg()
	page := rule.InitPage
	go func() {
		for {
			ch <- rule.RequestRss(initMsg)
			if !rule.CanPage {
				close(ch)
				break
			}
			page++
			rule.GeneratePageMsg(page, initMsg)
		}
	}()
	return ch
}

func (rule *XmlRule) GenerateMsgChan() chan *ParseResult {
	ch := make(chan *ParseResult)
	go func() {
		parentCh := rule.GenerateParentchan()
		for {
			parentNode, ok := <-parentCh
			if !ok {
				close(ch)
				break
			}
			for _, val := range parentNode {
				ch <- rule.GetResult(val)
			}
		}
		fmt.Println("generate msg chan stop")
	}()
	return ch
}
