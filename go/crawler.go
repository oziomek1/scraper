package main

import (
	"fmt"
	"golang.org/x/net/html"
)

// -------------------------------------
// Structure with info about each offers page
// -------------------------------------
type PageData struct {
	pageNumber int
	links []string
	nextPageUrl string
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
// Find links for offers and print them
// -------------------------------------
func getAndShowLinks(htmlTag string, page *html.Node, iteration int) (links[] string) {
	links = getLinks(htmlTag, page, links)
	for idx, link := range links {
		fmt.Println(idx + 32*iteration + 1, " -> ", link)
	}
	return links
}

// -------------------------------------
// Find next page button if exists
// -------------------------------------
func getAndShowNextPageUrl(page *html.Node) (nextPageUrl string) {
	nextPageUrl = nextPageLink(page)

	if nextPageUrl != "" {
		fmt.Println("Link to the next page:", nextPageUrl)
	}
	return nextPageUrl
}

func visitOffers(data PageData, offers []Offer) {
	for _, link := range data.links {
		go visitOffer(link, offers)
	}
}

// -------------------------------------
// Main loop, iterate whenever next page button exists
// -------------------------------------
func urlLinkCrawl(htmlTag string, url string, pageData []PageData, offers []Offer, iteration int) {
	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return
	}
	fmt.Println("\nCurrent page title: ", pageTitle(pageContent))

	currentData := PageData{iteration + 1,  getAndShowLinks(htmlTag, pageContent, iteration), getAndShowNextPageUrl(pageContent)}
	pageData = append(pageData, currentData)

	visitOffers(currentData, offers)

	for _, o := range offers {
		fmt.Println("PRINTUJ", o.url)
	}

	if currentData.nextPageUrl != "" {
		iteration += 1
		urlLinkCrawl(htmlTag, currentData.nextPageUrl, pageData, offers, iteration)
	} else {
		fmt.Println("The end of offers")
	}
}
