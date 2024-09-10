package main

import (
	"errors"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)

	if err != nil {
		return "", errors.New("error while getting HTML from website")
	}
	if res.StatusCode > 399 {
		return "", errors.New("response status code is error-level(400+)")
	}
	if res.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("invalid content-type")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("error while processing response")
	}

	return string(body), nil
}
