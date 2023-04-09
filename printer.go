package curlcolor

import (
	"github.com/alecthomas/chroma/v2/quick"
	"io"
)

type ColorPrinter struct {
	OutWriter io.Writer
	ErrWriter io.Writer
	Formatter string
	Theme     string
}

func (c *ColorPrinter) Print(text []byte, lexer string) {
	if err := quick.Highlight(c.OutWriter, string(text), lexer, c.Formatter, c.Theme); err != nil {
		return
	}
}
