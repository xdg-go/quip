package quip_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xdg/maybe"
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

func TestStringsToWords(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	// Words from Lines
	lines := maybe.JustAoS(dataLines)
	words := lines.Bind(quip.StringsToWords)
	wfl, err := words.Unbox()
	is.Nil(err)
	is.Equal(wfl, dataWords)
}

func TestParser(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	// Empty input
	{
		q := newQuip("")
		ss, err := q.Lines().Unbox()
		is.Nil(err)
		is.Equal(ss, []string{})
	}

	// Lines from Parser
	{
		q := newQuip(dataInput)
		lines := q.Lines()
		ss, err := lines.Unbox()
		is.Nil(err)
		is.Equal(ss, dataLines)
	}

	// Words from Parser
	{
		q := newQuip(dataInput)
		wfp, err := q.Words().Unbox()
		is.Nil(err)
		is.Equal(wfp, dataWords)
	}

}
