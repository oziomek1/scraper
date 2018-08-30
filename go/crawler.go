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
// Find links for offers and print them
// -------------------------------------
func getAndShowLinks(htmlTag string, page *html.Node, iteration int, allLinks *[]string) (links[] string) {
	links = getLinks(htmlTag, page, links)
	*allLinks = append(*allLinks, links...)
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

func visitOffers(links []string, offers *[]Offer) {
	for _, link := range links {
		visitOffer(link, offers)
	}
}

// -------------------------------------
// Main loop
// - collect urls for offers
// - iterate whenever next page button exists
// -------------------------------------
func urlLinkCrawl(htmlTag string, url string, pageData *[]PageData, offers *[]Offer, allLinks *[]string, iteration int) {

	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return
	}

	currentData := PageData{iteration + 1,  getAndShowLinks(htmlTag, pageContent, iteration, allLinks), getAndShowNextPageUrl(pageContent)}
	*pageData = append(*pageData, currentData)x

	if currentData.nextPageUrl != "" {
		iteration += 1
		urlLinkCrawl(htmlTag, currentData.nextPageUrl, pageData, offers, allLinks, iteration)
	} else {
		fmt.Println("THIS IS THE END OF OFFERS")
	}
}
