package libs

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

func Crawler(url string) (string, error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("error")
		return "crawler error", err
	}
	nodes := htmlquery.FindOne(doc, `//article`)
	html := htmlquery.OutputHTML(nodes, true)
	return html, err
}
