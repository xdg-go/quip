package quip_test

import (
	"bytes"
	"testing"

	"github.com/xdg/quip"
	"github.com/xdg/testy"
)

func TestNew(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var b bytes.Buffer
	q := quip.New(&b)

	is.NotNil(q)
}
