package main

import (
	"golang.org/x/net/html"
	"strings"
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
func getElementsById(tag string, n *html.Node, elements []string) (*html.Node, []string) {
	var node *html.Node
	for _, a := range n.Attr {
		if a.Key == tag && a.Val != "a" {
			elements = append(elements, a.Val)
			return node, elements
		} else if a.Key == tag {
			node, elements = getElementsById(tag, n.FirstChild, elements)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node, elements = getElementsById(tag, c, elements)
	}
	return node, elements
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

// -------------------------------------
// Get value of each param
// -------------------------------------
func getParamValue(page *html.Node) (string) {
	searchable := page.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild

	if searchable.NextSibling == nil {
		return filterUnnecessaryChars(searchable.Data)
	} else {
		return filterUnnecessaryChars(searchable.NextSibling.FirstChild.Data)
	}
}

// -------------------------------------
// Get value of each param label
// -------------------------------------
func getParamLabel(page *html.Node) (string) {
	return filterUnnecessaryChars(page.FirstChild.NextSibling.FirstChild.Data)
}

// -------------------------------------
// Collect all the params with labels of the offer
// -------------------------------------
func getOfferParam(page *html.Node, values []string, labels []string) ([]string, []string) {
	var paramValue string
	var paramLabel string
	if page.Type == html.ElementNode && page.Data == "li"{
		if len(page.Attr) > 0 && page.Attr[0].Key == "class" && page.Attr[0].Val == "offer-params__item" {
			paramValue = getParamValue(page)
			paramLabel = getParamLabel(page)
			labels = append(labels, paramLabel)
			values = append(values, paramValue)
		}
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		values, labels = getOfferParam(c, values, labels)
		if paramValue != "" {
			break
		}
	}

	return values, labels
}

// -------------------------------------
// Prevent \n \t \r etc.
// -------------------------------------
func filterUnnecessaryChars(s string) string {
	s = strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r

		}
	}, s)
	s = strings.TrimSpace(s)
	return s
}