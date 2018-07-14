package main

import (
	"fmt"
	"golang.org/x/net/html"
)

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

func urlCrawl(links[] string) {
	for index, link := range links {
		fmt.Println(index, "->", link)
	}

	for index, link := range links {
		pageContent, err := parse(link)
		if err != nil {
			fmt.Printf("Error with %s %s", pageContent, err)
			return
		}
		fmt.Println(index, "->", pageTitle(pageContent))
	}
}

