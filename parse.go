package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link is a structure for html links
type Link struct {
	Href string
	Text string
}

// Parse accepts a reader containing html and returns the links from it
func Parse(r io.Reader) ([]Link, error) {

	// Parse reader as html
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// Get the links, fam
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// Search the attributes for an HREF
	for _, att := range n.Attr {
		if att.Key == "href" {
			ret.Href = att.Val
			break
		}
	}
	ret.Text = text(n)

	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	var ret []*html.Node

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}
