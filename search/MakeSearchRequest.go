package search

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"main/htmlParsing"
	"net/url"
)

func MakeSearchRequest(query string) error {

	if len([]byte(query)) == 0 {
		fmt.Println("[INFO] No search query provided")
		return nil
	}
	query = url.QueryEscape(url.QueryEscape(query))

	connection, error := tls.Dial("tcp", "html.duckduckgo.com:443", &tls.Config{})

	if error != nil {
		return fmt.Errorf("[ERROR] could not open a connection: \"%v\"\n", error)
	}
	defer connection.Close()

	var request = fmt.Sprintf("GET /html/?q=%s HTTP/1.0\r\n"+
		"Host: html.duckduckgo.com\r\n"+
		"accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\n"+
		"accept-language: en-US,en;q=0.9,ro;q=0.8\r\n"+
		"priority: u=0, i\r\n"+
		"sec-ch-ua: \"Chromium\";v=\"146\", \"Not-A.Brand\";v=\"24\", \"Google Chrome\";v=\"146\"\r\n"+
		"sec-ch-ua-mobile: ?1\r\n"+
		"sec-ch-ua-platform: \"Android\"\r\n"+
		"sec-fetch-dest: document\r\n"+
		"sec-fetch-mode: navigate\r\n"+
		"sec-fetch-site: none\r\n"+
		"sec-fetch-user: ?1\r\n"+
		"upgrade-insecure-requests: 1\r\n"+
		"user-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Mobile Safari/537.36\r\n"+
		"Connection: close\r\n\r\n", query)
	connection.Write([]byte(request))

	response, _, err := bufio.NewReader(connection).ReadLine()

	if err != nil {
		return fmt.Errorf("[ERROR] could not read the respose status: \"%v\"\n", error)
	}

	fmt.Println(string(response))

	htmlNodes, error := html.Parse(connection)

	if error != nil {
		return fmt.Errorf("[ERROR] could not PARSE the response: \"%v\"\n", error)
	}

	var searchResults = parsing.TraverseTree(htmlNodes)

	for i := range searchResults {
		result := searchResults[i]
		fmt.Printf("%d. Title: %s\nLink: %s\n\n", i+1, result.Title, result.Url)
	}

	return nil
}
