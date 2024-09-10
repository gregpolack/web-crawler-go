package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL := args[0]
	pages := make(map[string]int)

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	crawlPage(baseURL, baseURL, pages)

	for page, count := range pages {
		fmt.Println(page, count)
	}

}
