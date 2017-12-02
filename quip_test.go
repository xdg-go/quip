package quip_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xdg/quip"
	"github.com/xdg/testy"
)

var TestStrings = []string{"Line one", "Line two", "Line three"}
var TestInput = strings.Join(TestStrings, "\n")

func TestNew(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var b bytes.Buffer
	q := quip.New(&b)

	is.NotNil(q)
}

func TestLines(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	b := bytes.NewBufferString(TestInput)
	q := quip.New(b)

	ss, err := q.Lines().Unbox()
	is.Nil(err)
	is.Equal(ss, TestStrings)

}
