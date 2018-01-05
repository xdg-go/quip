package quip

import "github.com/xdg/maybe"

// WordsT is a variant type representing either a slice of strings or an
// error.
type WordsT struct {
	maybe.AoS
}
