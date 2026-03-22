package parsing

import (
	"golang.org/x/net/html"
)

func isSearchResultEl(attributes []html.Attribute) bool {
	for i := range attributes {
		var attr = attributes[i]

		if attr.Key == "class" && attr.Val == "result__a" {
			return true
		}
	}

	return false
}
