package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	maxPages           int
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	// Check the number of command-line arguments
	if len(os.Args) < 4 {
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	// If there's exactly one argument, it should be the BASE_URL
	rawBaseURL := os.Args[1]
	baseParsed, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		os.Exit(1)
	}
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a postive integer")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	pages := make(map[string]int)

	// Set up the config struct
	cfg := &config{
		maxPages:           maxPages,
		pages:              pages,
		baseURL:            baseParsed,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency), // Limit to 5 concurrent requests
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)

	cfg.wg.Wait()

	fmt.Println("crawling complete. Pages crawled:")
	for page, count := range cfg.pages {
		fmt.Printf("%s: %d\n", page, count)
	}

	fmt.Printf("max concurrency: %d\n", maxConcurrency)
	fmt.Printf("max pages: %d\n", maxPages)
	// html, err := getHTML(baseURL)
	// if err != nil {
	// 	fmt.Printf("error fetching HTML: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(html)
}
