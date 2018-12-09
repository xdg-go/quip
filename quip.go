package quip

import (
	"bufio"
	"io"
	"os"
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

// NewFile constructs a new quip.Parser from a filename
func NewFile(name string) (*Parser, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &Parser{r: r}, nil
}

// Lines returns an AoS representing either lines of the input or an error.
// Delimiting "\n" or "\r\n" are not included.  If the input reader is
// empty, a valid, but empty result is returned.
func (p Parser) Lines() maybe.AoS {
	slice := make([]string, 0)
	rdr := bufio.NewReader(p.r)

	// Check for empty/bad reader
	_, err := rdr.Peek(1)
	if err != nil {
		if err == io.EOF {
			return maybe.JustAoS(slice)
		}
		return maybe.ErrAoS(err)
	}

LOOP:
	for {
		s, err := rdr.ReadString('\n')
		switch {
		case err == nil:
			slice = append(slice, strings.TrimRight(s, "\r\n"))
		case err == io.EOF:
			slice = append(slice, strings.TrimRight(s, "\r\n"))
			break LOOP
		default:
			return maybe.ErrAoS(err)
		}
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
