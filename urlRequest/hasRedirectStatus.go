package urlrequest

import (
	"strings"
)

func hasRedirectStatus(headers []byte) bool {
	statusLine := strings.Split(string(headers), "\n")[0]
	return strings.Contains(statusLine, "30")
}
