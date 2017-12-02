package quip

import "io"

type Parser struct {
	r io.Reader
}

func New(r io.Reader) *Parser {
	return &Parser{r: r}
}
