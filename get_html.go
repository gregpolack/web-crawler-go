package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) error {
	res, err := http.Get(rawURL)

	if err != nil {
		return errors.New("error while getting HTML from website")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error while processing response")
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return errors.New("response status code is error-level(400+)")
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return errors.New("invalid content-type")
	}

	fmt.Println(body)

	return nil
}
