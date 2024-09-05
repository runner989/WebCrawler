package main

import (
	"net/url"
	"strings"
)

// normalizeURL accepts a URL string and returns a normalized version of the URL.
func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// normalizedURL := parsedURL.Hostname() + strings.TrimSuffix(parsedURL.Path, "/")
	// return normalizedURL, nil

	// Normalize the host (remove scheme, etc.)
	host := parsedURL.Hostname()
	port := parsedURL.Port()

	if port == "80" || port == "443" {
		// ignore default ports
		port = ""
	}
	if port != "" {
		host = host + ":" + port
	}

	path := parsedURL.Path
	if path != "/" {
		path = strings.TrimSuffix(path, "/")
	}

	return host + path, nil
}
