package curlcolor

import (
	"github.com/alecthomas/chroma/v2/quick"
	"os"
	"testing"
)

func TestColorPrinter_Print(t *testing.T) {
	text := `HTTP/2 200 OK
date: Sat, 08 Apr 2023 18:44:56 GMT
content-type: application/json
content-length: 256
server: gunicorn/19.9.0
access-control-allow-origin: *
access-control-allow-credentials: true

`
	quick.Highlight(os.Stdout, text, "http", "terminal256", "monokai")
}
