package main

import (
	"bufio"
	"net/http"
	"strings"
)

func ParseResponse(text, url string) (*http.Response, error) {
	// https://go-review.googlesource.com/c/go/+/259758
	// Note that the http.ParseHTTPVersion does not support HTTP/2 or HTTP/3 responses.
	if strings.HasPrefix(text, "HTTP/2 ") {
		text = "HTTP/2.0 " + text[7:]
	}
	if strings.HasPrefix(text, "HTTP/3 ") {
		text = "HTTP/3.0 " + text[7:]
	}
	r := bufio.NewReader(strings.NewReader(text))
	req, err := http.NewRequest("GET", url, r)
	if err != nil {
		return nil, err
	}
	resp, err := http.ReadResponse(r, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
