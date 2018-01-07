package quip

import (
	"bufio"
	"io"
	"strings"

	"github.com/xdg/maybe"
)

// Parser provides quick line, word, character, etc. parsing against an
// io.Reader.
type Parser struct {
	r io.Reader
}

// New constructs a new quip.Parser from an io.Reader.
func New(r io.Reader) *Parser {
	return &Parser{r: r}
}

// Lines returns an AoS representing either lines of the input or an error.
func (p Parser) Lines() maybe.AoS {
	slice := make([]string, 0)
	scanner := bufio.NewScanner(p.r)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return maybe.ErrAoS(err)
	}

	return maybe.JustAoS(slice)
}

// Words returns an AoS representing either all words from all lines of the
// input or an error.
func (p Parser) Words() maybe.AoS {
	return p.Lines().Bind(StringsToWords)
}

// StringsToWords is a Bind function to convert a slice of strings to an AoS
// of whitespace-separated words.
// XXX if xs is nil, needs to return ErrAoS
func StringsToWords(xs []string) maybe.AoS {
	words := make([]string, 0)
	for _, v := range xs {
		ys := strings.Fields(v)
		if len(ys) > 0 {
			words = append(words, ys...)
		}
	}
	return maybe.JustAoS(words)
}
