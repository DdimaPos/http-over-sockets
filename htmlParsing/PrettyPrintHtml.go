package parsing

import (
	"fmt"

	"jaytaylor.com/html2text"
)

func PrettyPrintHtml(html []byte) error {
	prettyHtml, err := html2text.FromString(string(html), html2text.Options{PrettyTables: true})

	if err != nil {
		return fmt.Errorf("Could not convert html to text")
	}

	fmt.Print(prettyHtml)
	return nil
}
