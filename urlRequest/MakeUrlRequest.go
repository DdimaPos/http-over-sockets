package urlrequest

import (
	"fmt"
	parsing "main/htmlParsing"
	"main/printing"
	"net/url"
)

func MakeUrlRequest(passedUrl string, depth int) error {
	if depth > 5 {
		return fmt.Errorf(printing.Red + "[ERROR] Maximum redirect depth exceeded\n" + printing.Reset)
	}

	urlObj, err := url.Parse(passedUrl)

	if len([]byte(passedUrl)) == 0 {
		return nil
	}

	if err != nil {
		return fmt.Errorf(printing.Red+"[ERROR] Could not parse the url: %s\n"+printing.Reset, err)
	}

	rawResponse, err := executeRequest(urlObj)

	if err != nil {
		return err
	}

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
