package main

import (
	"flag"
	"fmt"
	"main/search"
	urlrequest "main/urlRequest"
)

func main() {
	var searchQuery string
	var url string
	flag.StringVar(&searchQuery, "s", "", "can introduce here the search query")
	flag.StringVar(&url, "u", "", "specicify a url to make a request to that url")
	flag.Parse()

	search.MakeSearchRequest(searchQuery)
	error := urlrequest.MakeUrlRequest(url)

	if error != nil {
		fmt.Printf("Error occured\n%s", error)
	}
}
