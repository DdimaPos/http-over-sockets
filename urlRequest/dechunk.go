package urlrequest

import (
	"bytes"
	"fmt"
)

func dechunk(body []byte) []byte {
	var result []byte
	cursor := 0

	for cursor < len(body) {
		lineEnd := bytes.Index(body[cursor:], []byte("\r\n"))
		if lineEnd == -1 {
			break
		}

		var size int
		fmt.Sscanf(string(body[cursor:cursor+lineEnd]), "%x", &size)

		if size == 0 {
			break
		}

		cursor += lineEnd + 2
		result = append(result, body[cursor:cursor+size]...)
		cursor += size + 2
	}
	return result
}
