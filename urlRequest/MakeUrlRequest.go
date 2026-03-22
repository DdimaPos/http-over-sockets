package urlrequest

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"jaytaylor.com/html2text"
	"net"
	"net/url"
)

func MakeUrlRequest(passedUrl string) error {
	var fullPath string
	var err error

	if len([]byte(passedUrl)) == 0 {
		fmt.Println("[INFO] No url for direct query provided")
		return nil
	}

	urlObj, err := url.Parse(passedUrl)

	if err != nil {
		return fmt.Errorf("[ERROR] Could not parse the url: %s\n", err)
	}

	hostname, pathname, queryParams := urlObj.Host, urlObj.Path, urlObj.RawQuery

	if len([]byte(hostname)) == 0 {
		return fmt.Errorf("[ERROR] Could not extract the hostname. Check the passed parameters")
	}

	if len([]byte(pathname)) == 0 {
		pathname = "/"
	}

	fullPath = pathname

	if len([]byte(queryParams)) > 0 {
		fullPath += "?" + queryParams
	}

	var conn net.Conn

	if urlObj.Scheme == "http" {
		conn, err = net.Dial("tcp", hostname+":80")
	} else {
		conn, err = tls.Dial("tcp", hostname+":443", &tls.Config{})
	}

	if err != nil {
		return fmt.Errorf("Could not open a tcp connection with %s", hostname)
	}

	defer conn.Close()

	request := fmt.Sprintf("GET %s HTTP/1.1\r\n"+
		"Host: %s\r\n"+
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\n"+
		"User-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Mobile Safari/537.36\r\n"+
		"Connection: close\r\n\r\n", fullPath, hostname)
	conn.Write([]byte(request))
	rawResponse, _ := io.ReadAll(conn)

	separator := "\r\n\r\n"
	splitIndex := bytes.Index(rawResponse, []byte(separator))

	if splitIndex == -1 {
		return fmt.Errorf("Could not separate headers from the response")
	}

	headers := rawResponse[:splitIndex]
	body := rawResponse[splitIndex+4:]

	// if hasRedirectStatus(headers) {
	//
	// 	newLocation := getRedirectLocation(headers)
	//
	// 	MakeUrlRequest(urlObj.Scheme + "://" + newLocation + fullPath)
	//
	// 	return nil
	// }

	fmt.Print("Headers:\n" + string(headers))

	prettyHtml, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})

	if err != nil {
		return fmt.Errorf("Could not pretty print the html")
	}
	fmt.Print(prettyHtml)

	return nil
}
