package utils

import (
	"io"

	"golang.org/x/net/html"
)

func ParseLink(htmlReader io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	for n := range doc.Descendants() {
		var link Link
		if n.Type == html.TextNode && n.Parent.Data == "a" {
			link.Text = n.Data
			for _, a := range n.Parent.Attr {
				if a.Key == "href" {
					link.Href = a.Val
				}
			}
		}
		if link.Href != "" && link.Text != "" {
			links = append(links, link)
		}

	}
	return links, nil
}
