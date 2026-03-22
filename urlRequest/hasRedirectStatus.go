package urlrequest

import (
	"fmt"
	"strings"
)

func hasRedirectStatus(headers []byte) bool {
	statusLine := strings.Split(string(headers), "\n")[0]
	fmt.Println(strings.Contains(statusLine, "30"))
	return strings.Contains(statusLine, "30")
}
