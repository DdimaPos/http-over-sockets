package search

import (
	"fmt"
	"golang.org/x/net/html"
)

func isResultsClassName(attributes []html.Attribute) bool {
	for i := range attributes {
		var attr = attributes[i]

		if attr.Key == "class" && attr.Val == "result__a" {

			fmt.Printf("FOUND %s, %s\n", attr.Val, attr.Key)
			return true
		}
	}

	return false
}
