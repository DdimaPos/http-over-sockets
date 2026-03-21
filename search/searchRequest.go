package search

// for URL navigation: if response is 302, check the Location header and make a new request

import (
	"crypto/tls"
	"fmt"

	"golang.org/x/net/html"
)

type SearchResult struct {
	url   string
	title string
}

func MakeSearchRequest(query string) error {

	if len([]byte(query)) == 0 {
		fmt.Println("[INFO] No search query provided")
		return nil
	}

	connection, error := tls.Dial("tcp", "html.duckduckgo.com:443", &tls.Config{})

	if error != nil {
		return fmt.Errorf("[ERROR] could not open a connection: \"%v\"\n", error)
	}
	defer connection.Close()

	request := fmt.Sprintf("GET /html/?q=%s HTTP/1.0\r\n"+
		"Host: html.duckduckgo.com\r\n"+
		"accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\n"+
		"accept-language: en-US,en;q=0.9,ro;q=0.8\r\n"+
		"priority: u=0, i\r\n"+
		"sec-ch-ua: \"Chromium\";v=\"146\", \"Not-A.Brand\";v=\"24\", \"Google Chrome\";v=\"146\"\r\n"+
		"sec-ch-ua-mobile: ?1\r\n"+
		"sec-ch-ua-platform: \"Android\"\r\n"+
		"sec-fetch-dest: document\r\n"+
		"sec-fetch-mode: navigate"+
		"sec-fetch-site: none\r\n"+
		"sec-fetch-user: ?1\r\n"+
		"upgrade-insecure-requests: 1\r\n"+
		"user-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Mobile Safari/537.36\r\n"+
		"Connection: close\r\n\r\n", query)
	connection.Write([]byte(request))

	htmlNodes, error := html.Parse(connection)

	if error != nil {
		return fmt.Errorf("[ERROR] could not PARSE the response: \"%v\"\n", error)
	}

	traverse(htmlNodes)

	// pick a list of all elements that have data-test-id="result-title-a"
	// go through first ten and extract the href and text inside the tag
	// create a struct var were I will save the url and title
	// put the struct into the list

	// in a loop print the list
	// accept the

	// tpe := htmlNodes.LastChild.Type

	// fmt.Println(string(htmlNodes))
	return nil
}
