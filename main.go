package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"sync"
)

type pageInfo struct {
	URL   string
	Count int
}

type config struct {
	maxPages           int
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func printReport(pages map[string]int, baseURL string) {
	var pageList []pageInfo
	for url, count := range pages {
		pageList = append(pageList, pageInfo{URL: url, Count: count})
	}

	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].Count == pageList[j].Count {
			return pageList[i].URL < pageList[j].URL
		}
		return pageList[i].Count > pageList[j].Count
	})

	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	for _, page := range pageList {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
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

	// fmt.Println("crawling complete. Pages crawled:")
	// for page, count := range cfg.pages {
	// 	fmt.Printf("%s: %d\n", page, count)
	// }

	// Print the report after crawling is complete
	printReport(cfg.pages, rawBaseURL)
}
