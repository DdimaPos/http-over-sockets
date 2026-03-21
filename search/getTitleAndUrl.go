package search

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getTitleAndUrl(n *html.Node) (title, url string) {
	var rawHref string

	for i := range n.Attr {
		attr := n.Attr[i]
		if attr.Key == "href" {
			rawHref = strings.Split(attr.Val, "uddg=")[1]
		}
	}

	url = cleanDuckDuckUrl(rawHref)

	return "adf", "adf"
}

func cleanDuckDuckUrl(rawUrl string) string {

	cleanUrl, _ := url.QueryUnescape(rawUrl)
	fmt.Printf(strings.Split(cleanUrl, "&")[0])

	return strings.Split(cleanUrl, "&")[0]
}
