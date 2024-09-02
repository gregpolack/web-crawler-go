package main

import (
	"errors"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	s, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.New("error normalizing the given URL")
	}
	if string(s.Path[len(s.Path)-1]) == "/" {
		s.Path = string(s.Path[:len(s.Path)-1])
	}
	return s.Host + s.Path, nil
}
