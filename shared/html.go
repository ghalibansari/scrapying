package shared

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func PrettyPrintHtml(text string) {
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		fmt.Println("could not parse properties table header html:", err)
		return
	}

	// Create a new bytes.Buffer and render the parsed HTML to it
	var b bytes.Buffer
	err = html.Render(&b, doc)
	if err != nil {
		fmt.Println("could not render properties table header html:", err)
		return
	}

	// Parse the rendered HTML as XML
	var prettyHTML bytes.Buffer
	decoder := xml.NewDecoder(strings.NewReader(b.String()))
	encoder := xml.NewEncoder(&prettyHTML)
	encoder.Indent("", "  ") // Set the indentation to 2 spaces

	// Copy the XML to the encoder to pretty print it
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			fmt.Println("could not encode token:", err)
			return
		}
	}

	// Flush the encoder to ensure all tokens are written
	err = encoder.Flush()
	if err != nil {
		fmt.Println("could not flush encoder:", err)
		return
	}

	// Print the pretty HTML
	fmt.Println(prettyHTML.String())
}
