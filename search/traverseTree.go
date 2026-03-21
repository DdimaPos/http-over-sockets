package search

import (
	"golang.org/x/net/html"
)

func traverse(n *html.Node) []SearchResult {
	var searchResults []SearchResult
	var newSearchResults []SearchResult

	if n.Type == html.ElementNode && isResultsClassName(n.Attr) {
		var title, url = getTitleAndUrl(n)
		searchResults = append(searchResults, SearchResult{title, url})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		newSearchResults = traverse(c)
	}

	return append(searchResults, newSearchResults...)
}
