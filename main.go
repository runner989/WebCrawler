package main

import (
	"fmt"
	"os"
)

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
	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	html, err := getHTML(baseURL)
	if err != nil {
		fmt.Printf("error fetching HTML: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(html)
}
