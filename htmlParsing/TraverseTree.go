package parsing

import (
	"main/structs"

	"golang.org/x/net/html"
)

var searchResults []structs.SearchResult

func TraverseTree(n *html.Node) []structs.SearchResult {

	if n.Type == html.ElementNode && isSearchResultEl(n.Attr) {
		var title, url = parseLinkTag(n)
		searchResults = append(searchResults, structs.SearchResult{Url: url, Title: title})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TraverseTree(c)
	}

	return searchResults
}
