package main

import (
	"net/http"
	"fmt"

	"golang.org/x/net/html"
)

// -------------------------------------
// Parse html page content to &html.Node type
// -------------------------------------
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
