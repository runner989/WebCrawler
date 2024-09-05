package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	// Traverse the HTML node tree to find all <a> tags
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Look for the "href" attribute
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := attr.Val
					parsedLink, err := url.Parse(link)
					if err == nil {
						// Resolve relative URLs to absolute URLs
						resolvedURL := baseURL.ResolveReference(parsedLink)
						urls = append(urls, resolvedURL.String())
					}
				}
			}
		}

		// Recursively traverse child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return urls, nil
}
