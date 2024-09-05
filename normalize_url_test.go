package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http scheme",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "root path",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev/",
		},
		{
			name:     "query string should be ignored",
			inputURL: "https://blog.boot.dev/path/?query=1",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "strip default port 80",
			inputURL: "http://blog.boot.dev:80/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "strip default port 443",
			inputURL: "https://blog.boot.dev:443/path",
			expected: "blog.boot.dev/path",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: unexpected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
