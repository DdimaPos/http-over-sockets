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
	var err error
	flag.StringVar(&searchQuery, "s", "", "can introduce here the search query")
	flag.StringVar(&url, "u", "", "specicify a url to make a request to that url")
	flag.Parse()

	err = search.MakeSearchRequest(searchQuery)
	if err != nil {
		fmt.Printf("Error occured\n%s", err)
	}

	err = urlrequest.MakeUrlRequest(url, 0)

	if err != nil {
		fmt.Printf("Error occured\n%s", err)
	}
}
