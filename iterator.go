package curlcolor

import (
	"github.com/alecthomas/chroma/v2"
	"strings"
)

func NewVerboseChannelIterator(source chan string) chroma.Iterator {
	var token *chroma.Token
	return func() chroma.Token {
		if token != nil {
			t := *token
			token = nil
			return t
		}
		for {
			line, ok := <-source
			if !ok {
				return chroma.EOF
			}

			switch true {
			case strings.HasPrefix(line, "curl:"):
				token = &chroma.Token{
					Type:  chroma.NameException,
					Value: line,
				}
			case strings.HasPrefix(line, "*"):
				return chroma.Token{
					Type:  chroma.CommentSingle,
					Value: line,
				}
			case strings.HasPrefix(line, ">"):
				token = &chroma.Token{
					Type:  chroma.CommentSingle,
					Value: line[1:],
				}
				return chroma.Token{
					Type:  chroma.NameException,
					Value: line[:1],
				}
			case strings.HasPrefix(line, "<"):
				token = &chroma.Token{
					Type:  chroma.CommentSingle,
					Value: line[1:],
				}
				return chroma.Token{
					Type:  chroma.Operator,
					Value: line[:1],
				}
			}
		}
	}
}
