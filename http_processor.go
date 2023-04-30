package curlcolor

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	url2 "net/url"
)

func SplitHeader(respBF *bufio.Reader) ([]byte, error) {
	var (
		header []byte
	)
	for {
		line, _, err := respBF.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(line) == 0 {
			header = append(header, '\n')
			break
		}
		header = append(header, line...)
		header = append(header, '\n')
	}
	return header, nil
}

func GetUrlScheme(url string) string {
	u, err := url2.Parse(url)
	if err != nil {
		return ""
	}
	return u.Scheme
}

func ParseAndPrintError(mgr *Manager, errReader io.Reader) error {
	bf := bufio.NewReader(errReader)
	ch := make(chan string)
	go func() {
		for {
			line, err := bf.ReadString('\n')
			if err != nil {
				close(ch)
				return
			}
			ch <- line
		}
	}()
	return mgr.Printer.VerboseChannelFormat(ch)
}

func ParseAndPrintOutput(mgr *Manager, output io.Reader) error {
	var (
		bf     = bufio.NewReader(output)
		header []byte
		err    error
		resp   *http.Response
	)
	for resp == nil ||
		(resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound && mgr.CurlParameter.GetBool("location")) {

		header, err = SplitHeader(bf)
		if err != nil {
			return err
		}
		if len(header) == 0 {
			break
		}
		if mgr.CurlParameter.GetBool("include") || mgr.CurlParameter.GetBool("head") {
			mgr.Printer.Highlight(header, "http")
		}
		resp, err = ParseResponse(header, mgr.CurlParameter.GetString("url"))
		if err != nil {
			return err
		}
	}
	body, err := io.ReadAll(bf)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	typ := GetFormatType(resp.Header.Get("Content-Type"), resp.Request.URL.Path)
	mgr.Printer.Highlight(body, typ)
	return nil
}

func ParseResponse(text []byte, url string) (*http.Response, error) {
	// https://go-review.googlesource.com/c/go/+/259758
	// Note that the http.ParseHTTPVersion does not support HTTP/2 or HTTP/3 responses.
	if bytes.HasPrefix(text, []byte("HTTP/2 ")) {
		text = append([]byte("HTTP/2.0 "), text[7:]...)
	}
	if bytes.HasPrefix(text, []byte("HTTP/3 ")) {
		text = append([]byte("HTTP/3.0 "), text[7:]...)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(text)), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
