package search

import (
	"fmt"
	"main/printing"
	"main/structs"
	urlrequest "main/urlRequest"
)

func pickSearchResult(searchResults []structs.SearchResult) error {
	var variant int
	fmt.Print("Which result you would like to access (enter the number): ")
	fmt.Scanln(&variant)

	if variant > len(searchResults) || variant <= 0 {
		fmt.Println(printing.Yellow + "[WARNING] Picked no variant or not from the list" + printing.Reset)
		return nil
	}

	return urlrequest.MakeUrlRequest(searchResults[variant-1].Url, 0)
}
