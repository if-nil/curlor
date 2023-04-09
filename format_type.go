package curlcolor

import (
	"path/filepath"
	"strings"
)

var mimeToFormatType = map[string]string{
	"application/json": "json",
	"application/xml":  "xml",
	"text/html":        "html",
	"text/xml":         "xml",
	"text/css":         "css",
}

func GetFormatType(contentType string, filename string) string {
	index := strings.Index(contentType, ";")
	if index > 0 {
		contentType = contentType[:index]
	}
	if mimeToFormatType[contentType] != "" {
		return mimeToFormatType[contentType]
	}
	extension := filepath.Ext(filename)
	if extension != "" {
		return extension[1:]
	}
	return ""
}
