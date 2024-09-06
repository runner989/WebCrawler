package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pageCount() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pageCount() >= cfg.maxPages {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing URL %s: %v\n", rawCurrentURL, err)
		return
	}
	// Ensure URL is on the same domain
	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		// fmt.Printf("skipping external link: %s\n", rawCurrentURL)
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	cfg.mu.Lock()
	if !cfg.addPageVisit(normalizedURL) {
		cfg.mu.Unlock()
		fmt.Printf("already crawled: %s\n", normalizedURL)
		return
	}
	cfg.mu.Unlock()

	fmt.Printf("crawling: %s\n", normalizedURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching HTML for %s: %v\n", rawCurrentURL, err)
		return
	}

	foundURLs, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, foundURL := range foundURLs {
		// cfg.mu.Lock()
		// if len(cfg.pages) >= cfg.maxPages {
		// 	cfg.mu.Unlock()
		// 	fmt.Println("max pages reached, stopping further crawling")
		// 	return
		// }
		// cfg.mu.Unlock()

		cfg.wg.Add(1)
		go cfg.crawlPage(foundURL)
	}
}
