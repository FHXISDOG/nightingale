package crawler

import (
	"fmt"
	"net/http"

	"github.com/antchfx/xmlquery"
)

const (
	// GET method
	GET = 1
	// POST method
	POST = 2
)

const (
	// URL xx
	URL = 1
	//BODY ss
	BODY = 2
)

const (
	//RSS the rss type
	RSS = 1
	//JSON type
	JSON = 2
)

var handleMap = make(map[int]func(rule *Rule) []*ParseResult)

func init() {
	handleMap[RSS] = handleRssRule
}

// ParseResult resulr of after analyse
type ParseResult struct {
	Title, Author, Description, Content, Link, Source string
	Date                                              string
}

// RuleMsg the Rule source
type SourceMsg struct {
	Description, Author string
	MsgType             int
}

// HTTPRule the http rule
type HTTPRule struct {
	URL           string
	RequestMethod int
	CanPage       bool
	InitPage      int
	CurrentPage   int
}

// MsgRule  the node of result represent
type MsgRule struct {
	ParentNode, TitleNode, DescriptionNode, ContentNode, LinkNode, DateNode string
}

// Rule of rss source
type Rule struct {
	*HTTPRule  `json:"HTTPRule"`
	*MsgRule   `json:"XmlMsgRule"`
	*SourceMsg `json:"SourceMsg"`
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

func (httpRule *HTTPRule) generateHttpInitmsg() *HttpInitMsg {
	return &HttpInitMsg{
		Url: httpRule.URL,
	}
}

func (httpRule *HTTPRule) generateNextpage(initMsg *HttpInitMsg, page int) {
	initMsg.Url = initMsg.Url
}

func (rule *Rule) generatePageMsg(page int, httpMsg *HttpInitMsg) {

}

func (httpRule *HTTPRule) requestMsg(initMsg *HttpInitMsg) *http.Response {
	url := initMsg.Url
	var resp *http.Response
	var err error
	switch httpRule.RequestMethod {
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
	return resp
}

func convertRssMsg2Result(rule *Rule, val *xmlquery.Node) *ParseResult {
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

func getRssParentNode(rule *Rule, resp *http.Response) []*xmlquery.Node {
	source := resp.Body
	defer source.Close()
	doc, _ := xmlquery.Parse(source)
	parenNode := xmlquery.Find(doc, rule.ParentNode)
	return parenNode
}

func handleRssRule(rule *Rule) []*ParseResult {
	httpReqMsg := rule.generateHttpInitmsg()
	resp := rule.requestMsg(httpReqMsg)
	parentNode := getRssParentNode(rule, resp)

	result := make([]*ParseResult, 0, len(parentNode))
	resCh := make(chan *ParseResult, len(parentNode))
	defer close(resCh)
	for _, val := range parentNode {
		go func(node *xmlquery.Node) {
			resCh <- convertRssMsg2Result(rule, node)
		}(val)
	}

	for val := range resCh {
		result = append(result, val)
		if len(result) == len(parentNode) {
			break
		}
	}

	return result
}

//GetResult get msg of rule
func (rule *Rule) GetResult() []*ParseResult {
	handleFunc, ok := handleMap[rule.MsgType]
	if ok {
		return handleFunc(rule)
	}
	return nil
}
