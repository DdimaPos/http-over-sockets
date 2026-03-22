package urlrequest

import "strings"

func getContentType(h []byte) string {
	var headers string = string(h)

	if strings.Contains(headers, "application/json") {
		return "json"
	}
	return "html"
}
