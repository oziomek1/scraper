package main

import (
	"fmt"

	"golang.org/x/net/html"
)
// -------------------------------------
// Imports for using goquery framework
//
//import (
//	"fmt"
//	"net/http"
//	"io/ioutil"
//	"github.com/PuerkitoBio/goquery"
//	"log"
//	"strings"
//)

// -------------------------------------

// Custom user agent.
const (
	userAgent = "Mozilla/5.0 (Macintosh; " +
	            "Intel Mac OS X 10_13_5) " +
	            "AppleWebKit/537.36 (KHTML, like Gecko) " +
	            "Chrome/67.0.3396.99 Safari/537.36"
)

// -------------------------------------

const (
	BASE_URL = "https://www.otomoto.pl/"

	// ---------------------------------
	// TYPE

	passenger = "osobowe/"
	delivery = "dostawcze/"
	motorcycle = "motocykle-i-quady/"
	truck = "ciezarowe/"
	construction = "maszyny-budowlane/"
	trailer = "przyczepy/"
	agro = "maszyny-rolnicze/"

	//----------------------------------
	// QUERIES

	QUERY_STRING_MARK = "?"
	QUERY_START = "search%5B"
	QUERY_JOIN = "&"

	//----------------------------------
	// SORT_TYPE

	time = "search%5Border%5D=created_at%3A"
	price = "search%5Border%5D=filter_float_price%3A"
	mileage = "search%5Border%5D=filter_float_mileage%3A"
	power = "search%5Border%5D=filter_float_engine_power%3A"

	//----------------------------------
	// SORT_TYPE_MODE

	asc = "asc"
	dsc = "desc"

	//----------------------------------
	// ENGINE_CAPACITY

	from = "search%5Bfilter_float_engine_capacity%3Afrom%5D="
	to = "search%5Bfilter_float_engine_capacity%3Ato%5D="

	//----------------------------------
	// YEARS

	YEAR_SINCE = "od-"
	YEAR_TO = "search%5Bfilter_float_year%3Ato%5D="
)
// -------------------------------------



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

func getOffersList(page *html.Node) string {
	var offersList string
	if page.Type == html.ElementNode && page.Data == "div" {
		fmt.Println(offersList)
	}
	return offersList
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
// Get next page link
// -------------------------------------
func nextPageLink(page *html.Node) (string) {
	var b html.Attribute
	var nextPage string
	for _, a := range page.Attr {
		if a.Key == "rel" && a.Val == "next" && b.Val != "" {
			nextPage = b.Val
			return nextPage
		}
		b = a
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		nextPage = nextPageLink(c)
	}
	return nextPage
}

func main() {

	make, model := "volkswagen/", "golf/"
	completeUrl := BASE_URL + passenger + make + model

	fmt.Println(completeUrl)

	page, err := parse(completeUrl)
	if err != nil {
		fmt.Printf("Error with %s %s", completeUrl, err)
		return
	}
	element, ok := getElementById("class", "offers list", page)

	fmt.Println("Page title: ", pageTitle(page))

	if !ok {
		fmt.Errorf("error finding element")
	} else {
		for _, a := range element.Attr {
			fmt.Println(a.Key, a.Val)
		}
	}

	var links []string
	nextPage := nextPageLink(page)
	fmt.Println(nextPage)
	tagWithLink := "data-href"
	links = getLinks(tagWithLink, page, links)

	urlCrawl(links)

}