package main

import (
	"path/filepath"
	"strings"
)

var mimeToFormatType = map[string]string{
	"application/json": "json",
	"application/xml":  "xml",
	"text/html":        "html",
	"text/plain":       "text",
	"text/xml":         "xml",
	"text/css":         "css",
}

func getFormatType(contentType string, filename string) string {
	index := strings.Index(contentType, ";")
	if index > 0 {
		contentType = contentType[:index]
	}
	if mimeToFormatType[contentType] != "" {
		return mimeToFormatType[contentType]
	}
	if filename != "" {
		extension := filepath.Ext(filename)
		if extension != "" {
			return extension[1:]
		}
	}
	return ""
}
