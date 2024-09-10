package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)

	if err != nil {
		return "", fmt.Errorf("error while getting HTML from website: %v", err)
	}
	if res.StatusCode > 399 {
		return "", fmt.Errorf("HTTP Error: %v", err)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error while processing response")
	}

	return string(body), nil
}
