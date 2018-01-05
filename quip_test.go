package quip_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xdg/quip"
	"github.com/xdg/testy"
)

var dataLines = []string{"Line one", "Line two", "Line three"}
var dataWords = []string{"Line", "one", "Line", "two", "Line", "three"}
var dataInput = strings.Join(dataLines, "\n")

func newQuip(s string) *quip.Parser {
	b := bytes.NewBufferString(s)
	return quip.New(b)
}

func TestEmpty(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	q := newQuip("")

	is.NotNil(q)
	ss, err := q.Lines().Unbox()
	is.Nil(err)
	is.Equal(ss, []string{})
}

func TestParser(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	q := newQuip(dataInput)

	// Lines from Parser
	lines := q.Lines()
	ss, err := lines.Unbox()
	is.Nil(err)
	is.Equal(ss, dataLines)

	// Words from Lines
	words := lines.Words()
	wfl, err := words.Unbox()
	is.Nil(err)
	is.Equal(wfl, dataWords)

	// Words from Parser
	q = newQuip(dataInput)
	wfp, err := q.Words().Unbox()
	is.Nil(err)
	is.Equal(wfp, dataWords)

}
