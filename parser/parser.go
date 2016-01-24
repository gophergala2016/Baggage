package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Parser struct {
	path string
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	Version string   `xml:"version,attr"`
	Title   string   `xml:"title"`
	Members Members  `xml:"members"`
	Items   Items    `xml:"items"`
}

type Members struct {
	XMLName xml.Name `xml:"members"`
	Member  []Member `xml:"member"`
}

type Member struct {
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name,attr"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"name,attr"`
	Owner   string   `xml:"owner,attr"`
}

func NewParser(p string) *Parser {
	return &Parser{path: p}
}

func (p *Parser) Parse() {

	f, err := os.Open(p.path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	prj := Project{}
	err = xml.Unmarshal(data, &prj)
	if err != nil {
		panic(err)
	}

	fmt.Println(prj)

}
