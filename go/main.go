package main

import (
	"fmt"
	"golang.org/x/net/html"
	time2 "time"
	"runtime"
	"runtime/debug"
	"os"
	"runtime/pprof"
	_ "net/http/pprof"
	_ "net/http"
	"net/http"
	"log"
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

	QUERY_MARK = "?"
	QUERY_START = "search%5B"
	QUERY_JOIN = "&"

	//----------------------------------
	// SORT_TYPE

	time = "order%5D=created_at%3A"
	price = "order%5D=filter_float_price%3A"
	mileage = "order%5D=filter_float_mileage%3A"
	power = "order%5D=filter_float_engine_power%3A"

	//----------------------------------
	// SORT_TYPE_MODE

	asc = "asc"
	dsc = "desc"

	//----------------------------------
	// ENGINE_CAPACITY

	engine_from = "filter_float_engine_capacity%3Afrom%5D="
	engine_to = "filter_float_engine_capacity%3Ato%5D="

	//----------------------------------
	// YEARS

	YEAR_SINCE = "od-"
	YEAR_TO = "filter_float_year%3Ato%5D="

	//----------------------------------
	// POWER

	power_from = "filter_float_engine_power%3Afrom%5D="
	power_to = "filter_float_engine_power%3Ato%5D="

)
// -------------------------------------

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
	runtime.GOMAXPROCS(32)

	start := time2.Now()

	var pageData []PageData
	var offers []Offer
	var allLinks []string

	make, model := "volkswagen/", "golf/"
	completeUrl := BASE_URL + passenger + make + model + QUERY_MARK + QUERY_START + power_from + "130"
	fmt.Println("\t\tStarting page url: ", completeUrl, "\n\n")

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	debug.PrintStack()

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	tagForLink := "data-href"
	urlLinkCrawl(tagForLink, completeUrl, &pageData, &offers, &allLinks, 0)

	visitOffers(allLinks, &offers)

	fmt.Println("number of offers:", len(offers))

	exportData(offers, "otomotoVWGolf4")

	fmt.Println(time2.Since(start))

}