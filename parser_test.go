package linkparser

import (
	"os"
	"strings"
	"testing"
)

func TestLinkParser(t *testing.T) {
	var tests = []struct {
		path   string
		expect []Link
	}{
		{
			path: "./examples/e1.html",
			expect: []Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				}},
		},
		{
			path: "./examples/e2.html",
			expect: []Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				}},
		},
		{
			path: "./examples/e3.html",
			expect: []Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				}},
		},
		{
			path: "./examples/e4.html",
			expect: []Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on",
				}},
		},
	}
	for i, test := range tests {
		file, err := os.ReadFile(test.path)
		if err != nil {
			t.Errorf("File read error\n\n%d\n\n", err)
			return
		}
		hmtl := string(file)
		got, err := ParseLinks(strings.NewReader(hmtl))
		if err != nil {
			t.Error(err)
		}
		if !compare(got, test.expect) {
			t.Errorf("Test Subject: %v\nGot	: %s\nWant	: %s", i, got, test.expect)
		}
	}

}
func compare(a []Link, b []Link) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
