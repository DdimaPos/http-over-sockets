package parsing

import (
	"fmt"
	"main/printing"

	"jaytaylor.com/html2text"
)

func PrettyPrintHtml(html []byte) error {
	prettyHtml, err := html2text.FromString(string(html), html2text.Options{PrettyTables: true})

	if err != nil {
		return fmt.Errorf(printing.Red + "[ERROR] Could not convert html to text" + printing.Reset)
	}

	fmt.Print(prettyHtml)
	return nil
}
