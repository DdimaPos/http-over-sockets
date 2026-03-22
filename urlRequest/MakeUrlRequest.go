package urlrequest

import (
	"fmt"
	parsing "main/htmlParsing"
	"net/url"
)

func MakeUrlRequest(passedUrl string, depth int) error {
	if depth > 5 {
		return fmt.Errorf("[ERROR] Maximum redirect depth exceeded")
	}

	urlObj, err := url.Parse(passedUrl)

	if len([]byte(passedUrl)) == 0 {
		fmt.Println("[INFO] No url for direct query provided")
		return nil
	}

	if err != nil {
		return fmt.Errorf("[ERROR] Could not parse the url: %s\n", err)
	}

	rawResponse, _ := executeRequest(urlObj)
	headers, body, err := splitHTTPResponse(rawResponse)

	if err != nil {
		return err
	}

	if hasRedirectStatus(headers) {
		newLocation := resolveRedirectUrl(urlObj, getRedirectLocation(headers))
		return MakeUrlRequest(newLocation, depth+1)
	}

	return parsing.PrettyPrintHtml(body)
}
