package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Fetch the webpage
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for an error-level status code (400+)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	// Ensure the content-type is text/html
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", errors.New("response content-type is not text/html")
	}

	// Read the HTML content
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBytes), nil
}
