package run

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"matt-rickard.com/virgo/pkg/virgo"
)

func TestParseVirgo(t *testing.T) {
	f, err := os.Open("testdata/build.vgo")
	if err != nil {
		t.Fatal("opening build.vgo")
	}
	defer f.Close()
	if err := ParseVirgo(f, os.Stdout); err != nil {
		t.Fatal("parsing")
	}
}

func ParseVirgo(in io.Reader, out io.Writer) error {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return errors.Wrap(err, "input")
	}
	g, err := virgo.Parse(buf)
	if err != nil {
		return errors.Wrap(err, "parsing")
	}
	sortedV, err := virgo.TopSort(g)
	if err != nil {
		return errors.Wrap(err, "sorting")
	}
	if _, err := io.WriteString(out, strings.Join(sortedV, "\n")); err != nil {
		return errors.Wrap(err, "output")
	}
	return nil
}
