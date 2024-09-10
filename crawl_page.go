package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Compare the domain name of the base and current URL
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
	}
	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	// Normalize the current URL for map logic
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL]++
		return
	} else {
		pages[normalizedCurrentURL] = 1
	}

	// Get the HTML of the current URL and get a slice of URLs inside it
	currentURLHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Getting HTML from: %s\n", rawCurrentURL)
	currentPageURLs, err := getURLsFromHTML(currentURLHTML, rawBaseURL)
	if err != nil {
		fmt.Println(err)
	}
	// Recursively call the function on each URL
	for _, nextURL := range currentPageURLs {
		crawlPage(rawBaseURL, nextURL, pages)

	}

}
