package crawler

import "io"

type ParseResult struct {
	Title, Author, Description, Content, Link, Source string
	Date                                              string
}

type XmlRule struct {
	ParentNode, TitleNode, DescriptionNode, ContentNode, LineNode, DateNode string
}

// parse xml
func xml_parser(source io.Reader, rule XmlRule) ParseResult {
	var result ParseResult

	return result
}
