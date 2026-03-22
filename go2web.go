package main

import (
	"flag"
	"main/search"
)

func main() {
	var searchQuery string
	var url string
	flag.StringVar(&searchQuery, "s", "", "can introduce here the search query")
	flag.StringVar(&url, "u", "", "specicify a url to make a request to that url")
	flag.Parse()

	search.MakeSearchRequest(string(searchQuery))
}
