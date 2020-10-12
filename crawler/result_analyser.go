package crawler

import (
	"fmt"
	"io"

	"github.com/antchfx/xmlquery"
)

type ParseResult struct {
	Title, Author, Description, Content, Link, Source string
	Date                                              string
}

type XmlRule struct {
	ParentNode, TitleNode, DescriptionNode, ContentNode, LinkNode, DateNode string
}

// parse xml
func xml_parser(source io.Reader, rule XmlRule) []*ParseResult {
	result := make([]*ParseResult, 100)
	doc, err := xmlquery.Parse(source)
	if err != nil {
		fmt.Println("parse xml error", err)
	}
	parentNode := xmlquery.find(doc, rule.ParentNode)
	for _, val := range parentNode {
		title := val.SelectElement(rule.TitleNode).InnerText()
		description := val.SelectElement(rule.DescriptionNode).InnerText()
		content := val.SelectElement(rule.ContentNode).InnerText()
		link := val.SelectElement(rule.LinkNode).InnerText()
		date := val.SelectElement(rule.DateNode).InnerText()

		result = result.append(&ParseResult{
			Title:       title,
			Content:     content,
			Description: description,
			Date:        date,
			Link:        link,
		})
	}
	return result
}
