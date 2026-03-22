package urlrequest

import (
	"fmt"
	"iter"
	"main/printing"
	"strings"
)

func getRedirectLocation(headers []byte) string {
	var lines iter.Seq[string] = strings.Lines(string(headers))

	for line := range lines {
		if strings.Contains(line, "ocation") {
			locationRaw := strings.Split(line, " ")[1]
			suffix, _ := strings.CutSuffix(locationRaw, "\r\n")
			fmt.Printf(printing.Blue+"[INFO] Found redirect Header: %s\n"+printing.Reset, suffix)
			return suffix
		}
	}

	return ""
}
