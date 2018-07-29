package main

import (
	"golang.org/x/net/html"
)

// -------------------------------------
// Example usage:
// <input tag="my_id_name" value="input_value"/>
// tag = "tag"
// id = "my_id_name"
// getElementById("tag", "my_id_name", pageNode)
// -------------------------------------
func getElementById(tag string, id string, n* html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == tag && a.Val == id {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling{
		if element, ok = getElementById(tag, id, c); ok {
			return
		}
	}
	return
}

// -------------------------------------
// Get elements by ID inside table
// -------------------------------------
func getElementsById(tag string, id string, n *html.Node, elements []string) (element *html.Node, elems []string, ok bool) {
	for _, a := range n.Attr {
		if a.Key == tag && a.Val == id {
			elements = append(elements, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, elements, ok = getElementsById(tag, id, c, elements); ok {
			return
		}
	}
	return
}

// -------------------------------------
// Get element inside xml node
// -------------------------------------
func getElementByTag(tag string, page *html.Node) (string) {
	var element string
	for _, a := range page.Attr {
		if a.Key == tag {
			element = a.Val
			return element
		}
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		element = getElementByTag(tag, c)
		if element != "" {
			break
		}
	}
	return element
}

// -------------------------------------
// Get links within particular xml node
// -------------------------------------
func getLinks(tag string, page *html.Node, links []string) ([]string) {
	for _, a := range page.Attr {
		if a.Key == tag {
			links = append(links, a.Val)
		}
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		links = getLinks(tag, c, links)
	}
	return links
}


// -------------------------------------
// Get page element value by tag name
// -------------------------------------
func pageElementValue(tag string, page *html.Node) (string) {
	var value string
	if page.Type == html.ElementNode && page.Data == tag {
		value = page.FirstChild.Data
		return value
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		value = pageTitle(c)
		if value != "" {
			break
		}
	}
	return value
}

// -------------------------------------
// Get page title
// -------------------------------------
func pageTitle(page *html.Node) (string) {
	var title string
	if page.Type == html.ElementNode && page.Data == "title" {
		title = page.FirstChild.Data
		return title
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// -------------------------------------
// Get next page link (button id)
// -------------------------------------
func nextPageLink(page *html.Node) (nextPage string) {
	var b html.Attribute
	for _, a := range page.Attr {
		if a.Key == "rel" && a.Val == "next" && b.Val != "" {
			nextPage = b.Val
			return nextPage
		}
		b = a
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		if nextPageLink(c) != "" {
			nextPage = nextPageLink(c)
		}
	}
	return nextPage
}