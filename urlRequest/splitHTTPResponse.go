package urlrequest

import (
	"bytes"
	"fmt"
)

func splitHTTPResponse(rawResponse []byte) (headers, body []byte, err error) {
	separator := "\r\n\r\n"
	splitIndex := bytes.Index(rawResponse, []byte(separator))

	if splitIndex == -1 {
		return headers, body, fmt.Errorf("Could not separate headers from the response")
	}

	headers = rawResponse[:splitIndex]
	body = rawResponse[splitIndex+4:]
	return headers, body, nil
}
