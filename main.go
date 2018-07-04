package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"log"
	"fmt"
	"io"
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

func fetchUrl(url string) (*goquery.Document, *goquery.Document) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	html, err := goquery.NewDocumentFromReader(io.Reader(resp.Body))
	if err != nil {
		panic(err)
	}

	text, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return html, text
}

func processElement(article_list *goquery.Selection) {
	href, present := article_list.Attr("href")
	if present {
		fmt.Println(href)
	}
}

func articles(html goquery.Document) {

	article_list := html.Find(".offers.list")
	processElement(article_list)
}

func main() {

	make, model := "volkswagen/", "golf/"
	completeUrl := BASE_URL + passenger + make + model

	fmt.Printf("HTML code of %s", completeUrl)

	html, _ := fetchUrl(completeUrl)
	articles(*html)

	//fmt.Printf("%s", html)

}
//
//const html = `
//<div class="container">
//    <div class="row">
//      <div class="col-lg-8">
//        <p align="justify"><b>Name</b>Priyaka</p>
//        <p align="justify"><b>Surname</b>Patil</p>
//        <p align="justify"><b>Adress</b><br>India,Kolhapur</p>
//        <p align="justify"><b>Hobbies&nbsp;</b><br>Playing</p>
//        <p align="justify"><b>Eduction</b><br>12th</p>
//        <p align="justify"><b>School</b><br>New Highschool</p>
//       </div>
//    </div>
//</div>
//`
//
//func main() {
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err != nil {
//		panic(err)
//	}
//
//	doc.Find(".container").Find("[align=\"justify\"]").Each(func(_ int, s *goquery.Selection) {
//		prefix := s.Find("b").Text()
//		result := strings.TrimPrefix(s.Text(), prefix)
//		println(result)
//	})
//}