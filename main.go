package main

import (
	"net/http"
	"fmt"

	"golang.org/x/net/html"
)

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

//func fetchUrl(url string) (*goquery.Document, *goquery.Document, *http.Response) {
//
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//
//	defer resp.Body.Close()
//	html, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))
//	if err != nil {
//		panic(err)
//	}
//
//	text, err := goquery.NewDocument(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return html, text, resp
//}
//
//func processElement(article_list *goquery.Selection) {
//	href, present := article_list.Attr("href")
//	if present {
//		fmt.Println(href)
//	}
//}

//func articles(html goquery.Document) {
//
//	article_list := html.Find(".offers.list")
//	processElement(article_list)
//}

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

func parse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot get page")
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse page")
	}

	return body, err
}

func getLinks(page *html.Node, links []string) ([]string) {
	var b html.Attribute
	for _, a := range page.Attr {
		if a.Key == "data-href" {
			links = append(links, a.Val)
		}
		if a.Key == "rel" && a.Val == "next" && b.Val != "" {
			fmt.Printf("Next page %s\n",  b.Val)
		}
		b = a
	}
	for c := page.FirstChild; c != nil; c = c.NextSibling {
		links = getLinks(c, links)
	}
	return links
}

func main() {

	make, model := "volkswagen/", "golf/"
	completeUrl := BASE_URL + passenger + make + model

	fmt.Println(completeUrl)

	//_, _, resp := fetchUrl(completeUrl)
	//articles(*html)
	//
	//resp, err := http.Get(completeUrl)
	//if err != nil {
	//	return
	//}
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Printf("%s", html)

	page, err := parse(completeUrl)
	if err != nil {
		fmt.Printf("Error with %s %s", completeUrl, err)
		return
	}
	var links []string
	links = getLinks(page, links)
	for index, link := range links {
		fmt.Println(index, "->", link)
	}

	element, ok := getElementById("class", "offers list", page)

	if !ok {
		fmt.Errorf("error finding element")
	} else {
		for _, a := range element.Attr {
			fmt.Println(a.Key, a.Val)
		}
	}

}