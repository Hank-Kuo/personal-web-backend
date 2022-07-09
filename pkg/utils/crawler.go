package utils

import (
	"github.com/antchfx/htmlquery"
)

func Crawler(url string) (string, error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", err
	}
	nodes := htmlquery.FindOne(doc, `//article`)
	html := htmlquery.OutputHTML(nodes, true)
	return html, err
}
