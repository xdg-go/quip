package quip

import (
	"strings"

	"github.com/xdg/maybe"
)

// LinesT is a variant type representing either a slice of strings or an
// error.
type LinesT struct {
	maybe.AoS
}

// Words returns a WordsT representing all words or a propogated error.
func (lt LinesT) Words() WordsT {
	return WordsT{AoS: lt.Bind(linesToWords)}
}

func linesToWords(xs []string) maybe.AoS {
	words := make([]string, 0)
	for _, v := range xs {
		ys := strings.Fields(v)
		if len(ys) > 0 {
			words = append(words, ys...)
		}
	}
	return maybe.JustAoS(words)
}
