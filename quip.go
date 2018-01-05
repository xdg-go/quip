package quip

import (
	"bufio"
	"io"

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

// Lines returns a LinesT representing either lines of the input or an error.
func (p Parser) Lines() LinesT {
	slice := make([]string, 0)
	scanner := bufio.NewScanner(p.r)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return LinesT{AoS: maybe.ErrAoS(err)}
	}

	return LinesT{AoS: maybe.JustAoS(slice)}
}

// Words returns a WordsT representing either all words from all lines of the
// input or an error.
func (p Parser) Words() WordsT {
	return p.Lines().Words()
}
