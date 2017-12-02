package quip

import (
	"bufio"
	"io"
)

// Parser provides quick line, word, character, etc. parsing against an
// io.Reader.
type Parser struct {
	r io.Reader
}

// LinesT is a variant type representing either a slice of strings or an error
type LinesT struct {
	ss  []string
	err error
}

// New construct a new quip.Parser from an io.Reader
func New(r io.Reader) *Parser {
	return &Parser{r: r}
}

// Lines returns a LinesT representing either lines of the input or an error.
func (p Parser) Lines() LinesT {
	ss := make([]string, 0)
	scanner := bufio.NewScanner(p.r)

	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return LinesT{err: err}
	}

	return LinesT{ss: ss}
}

// Unbox returns the underlying slice of strings and error.  One or the other
// will be nil.
func (lt LinesT) Unbox() ([]string, error) {
	return lt.ss, lt.err
}
