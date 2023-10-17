package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	var linkNodes []*html.Node

	findLinkNodes(doc, &linkNodes)
	if err != nil {
		panic(err)
	}
	for _, v := range linkNodes {
		ok, l := makeLinkFromNode(v)
		if ok {
			links = append(links, l)
		}
	}

	return links, nil
}

func findLinkNodes(n *html.Node, l *[]*html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinkNodes(c, l)

	}
	if n.Type == html.ElementNode && n.Data == "a" {
		*l = append(*l, n)
	}

}

func makeLinkFromNode(n *html.Node) (bool, Link) {
	for _, v := range n.Attr {
		if v.Key == "href" {

			return true, Link{
				Href: v.Val,
				Text: strings.Trim(n.FirstChild.Data, "\n \t"),
			}
		}
	}
	return false, Link{}

}
