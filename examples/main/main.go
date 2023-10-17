package main

import (
	"fmt"
	"os"
	"strings"

	lp "github.com/kmlgkcy/link-parser"
)

func main() {
	file, err := os.ReadFile("./examples/e4.html")
	if err != nil {
		fmt.Printf("File read error\n\n%d\n\n", err)
		return
	}

	hmtl := string(file)
	got, err := lp.ParseLinks(strings.NewReader(hmtl))
	if err != nil {
		panic(err)
	}
	for _, v := range got {
		fmt.Printf(`
{
	href: %s
	text: %s
}
		`, v.Href, v.Text)
	}
}
