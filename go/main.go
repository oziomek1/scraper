package main

import (
	"fmt"

	"golang.org/x/net/html"
	"runtime"
)

// -------------------------------------
// Imports if using goquery framework
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

	engine_from = "search%5Bfilter_float_engine_capacity%3Afrom%5D="
	engine_to = "search%5Bfilter_float_engine_capacity%3Ato%5D="

	//----------------------------------
	// YEARS

	YEAR_SINCE = "od-"
	YEAR_TO = "search%5Bfilter_float_year%3Ato%5D="

	//----------------------------------
	// POWER

	power_from = "search%5Bfilter_float_engine_power%3Afrom%5D="
	power_to = "search%5Bfilter_float_engine_power%3Ato%5D="

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


func main() {
	// ---------------------------------
	// Define max number of goroutines (lightweight threads)
	// ---------------------------------
	runtime.GOMAXPROCS(5)

	var pageData []PageData
	var offers []Offer

	make, model := "volkswagen/", "golf/"
	completeUrl := BASE_URL + passenger + make + model

	fmt.Println("\t\tStarting page url: ", completeUrl, "\n\n\n\n")

	tagForLink := "data-href"
	urlLinkCrawl(tagForLink, completeUrl, pageData, offers,0)

}