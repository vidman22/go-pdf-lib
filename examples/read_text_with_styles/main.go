package main

import (
	"fmt"

	"github.com/superpowerdotcom/go-pdf-lib"
)

func main() {
	f, r, err := pdf.Open("/Users/johnvidmar/Desktop/Code/go-pdf-lib/examples/read_text_with_styles/pdf_test.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sentences, err := r.GetStyledTexts()
	if err != nil {
		panic(err)
	}

	// Print all sentences
	for _, sentence := range sentences {
		fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n",
			sentence.Font,
			sentence.FontSize,
			sentence.X,
			sentence.Y,
			sentence.S)
	}
}
