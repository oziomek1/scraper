package main

import (
	"fmt"
	"golang.org/x/net/html"
	"sync"
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
func getAndShowLinks(htmlTag string, page *html.Node, allLinks *[]string) (links[] string) {
	links = getLinks(htmlTag, page, links)
	*allLinks = append(*allLinks, links...)
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
	var unReadedUrls []string
	miniBatchSize := 100
	for i := 0; i < len(links); i += miniBatchSize {
		miniBatchLinks := links[i:i+miniBatchSize]
		var wg = sync.WaitGroup{}
		wg.Add(len(miniBatchLinks))
		for _, link := range miniBatchLinks {
			go visitOffer(link, offers, &unReadedUrls, &wg)
		}
		wg.Wait()
	}

	// -------------------------------------
	// Used for crawl through the offers with no collected parameters.
	// Probably possible to remove, the problem with lack of these parameters
	// is related to otomoto itself (Access denied)
	// -------------------------------------
	for _, link := range unReadedUrls {
		visitOffer(link, offers, &unReadedUrls, nil)
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
		fmt.Println("Error with", pageContent, err)
		return
	}

	currentData := PageData{iteration + 1,  getAndShowLinks(htmlTag, pageContent, allLinks), getAndShowNextPageUrl(pageContent)}
	*pageData = append(*pageData, currentData)

	if currentData.nextPageUrl != "" {
		iteration += 1
		urlLinkCrawl(htmlTag, currentData.nextPageUrl, pageData, offers, allLinks, iteration)
	} else {
		fmt.Println("THIS IS THE END OF OFFERS")
	}
}
