package crawler

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"testing"
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
	elementRule := []string{"channel", "item"}
	fmt.Println(elementRule)
	result := getXmlElement(resp.Body, "channel")
	for _, value := range result {
		fmt.Println(value)
	}
}

func compareSliceEqual(a,b []string)  bool{
	for aIdx,aValue := range a{
		if b[aIdx] == nil || b[aIdx] != aValue{
			return false 
			break	
		}
	}
	return true
}

func strSliceContain(s []string,str string) bool{
	for _,value := range s{
		if str == value{
			return true
		}
	}
	return false
}

funcxml2Map(source io.Reader, rule []string) []map[string]string {
	result := make([]map[string]string, 10)
	decoder := xml.NewDecoder(source)
	parent := make([]string, 10)
	for {

		if compareSliceEqual(rule,parent){
			child := make(map[string]string)
			result = append(result, child)
		}
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		// 开始判断token类型
		switch element := token.(type) {
		case xml.StartElement:
			elementName := element.Name.Local
			if strSliceContain(rule,elementName){
				parent = append(parent, element.Name.Local)
			} 
		case xml.EndElement:
			if strSliceContain(rule,elementName){
				parent = append(parent, element.Name.Local)
			} 
		case xml.CharData:
			parent = parent[:len(parent)-1]
		default:
			// todosomething
		}

	}
	return result
}

func getXmlElementByRuleSlice(source io.Reader, elementRule []string) []Info {
	result := make([]Info, 100)
	decoder := xml.NewDecoder(source)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == elementRule[0] {
				element.End()
			}

		}
	}
	return result
}

func getXmlElement(source io.Reader, elementName string) []string {
	result := make([]string, 20)
	decoder := xml.NewDecoder(source)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == elementName {
				fmt.Println("enter the if statement")
				content, _ := decoder.Token()
				if contentStr, ok := content.(xml.CharData); ok {
					contentStr2 := string([]byte(contentStr))
					result = append(result, contentStr2)
				}

			}
		}
	}

	return result
}
