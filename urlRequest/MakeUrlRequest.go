package urlrequest

import (
	"fmt"
	// "net/tls"
	"net/url"
)

func MakeUrlRequest(passedUrl string) error {

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

	var fullPath string = pathname

	if len([]byte(queryParams)) > 0 {
		fullPath += "?" + queryParams
	}

	fmt.Println(hostname)
	fmt.Println(pathname)
	fmt.Println(queryParams)

	return nil
}
