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
// Main loop
// - collect urls for offers
// - iterate whenever next page button exists
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

	if currentData.nextPageUrl != "" {
		iteration += 1
		urlLinkCrawl(htmlTag, currentData.nextPageUrl, pageData, offers, iteration)
	} else {
		fmt.Println("THIS IS THE END OF OFFERS")
	}
}
