package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	// Check the number of command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// If there's exactly one argument, it should be the BASE_URL
	rawBaseURL := os.Args[1]
	baseParsed, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	pages := make(map[string]int)

	// Set up the config struct
	cfg := &config{
		pages:              pages,
		baseURL:            baseParsed,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5), // Limit to 5 concurrent requests
		wg:                 &sync.WaitGroup{},
	}
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)

	cfg.wg.Wait()

	fmt.Println("crawling complete. Pages crawled:")
	for page, count := range cfg.pages {
		fmt.Printf("%s: %d\n", page, count)
	}

	// html, err := getHTML(baseURL)
	// if err != nil {
	// 	fmt.Printf("error fetching HTML: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(html)
}
