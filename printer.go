package curlcolor

import (
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/alecthomas/chroma/v2/styles"
	"io"
)

type ColorPrinter struct {
	OutWriter io.Writer
	ErrWriter io.Writer
	Formatter string
	Theme     string
}

func (c *ColorPrinter) Highlight(text []byte, lexer string) {
	if err := quick.Highlight(c.OutWriter, string(text), lexer, c.Formatter, c.Theme); err != nil {
		return
	}
}

func (c *ColorPrinter) VerboseChannelFormat(text chan string) error {
	f := formatters.Get(c.Formatter)
	if f == nil {
		f = formatters.Fallback
	}
	s := styles.Get(c.Theme)
	if s == nil {
		s = styles.Fallback
	}
	it := NewVerboseChannelIterator(text)
	return f.Format(c.OutWriter, s, it)
}
