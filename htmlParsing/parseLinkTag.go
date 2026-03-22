package parsing

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func parseLinkTag(n *html.Node) (title, url string) {
	var rawHref string

	for i := range n.Attr {
		attr := n.Attr[i]
		if attr.Key == "href" {
			rawHref = strings.Split(attr.Val, "uddg=")[1]
			url = cleanDuckDuckUrl(rawHref)
			title = n.FirstChild.Data
		}
	}

	return title, url
}

func cleanDuckDuckUrl(rawUrl string) string {
	cleanUrl, _ := url.QueryUnescape(rawUrl)

	return strings.Split(cleanUrl, "&")[0]
}
