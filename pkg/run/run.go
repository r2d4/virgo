package run

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"matt-rickard.com/virgo/pkg/virgo"
)

func Run(in string) error {
	s := strings.Split(in, ":")
	var fname, root string
	fname = s[0]
	if len(s) == 2 {
		root = s[1]
	}
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}
	g, err := virgo.Parse(f)
	if err != nil {
		return errors.Wrap(err, "parsing virgo file")
	}

	nodes, err := virgo.DFS(g, root)
	if err != nil {
		return errors.Wrap(err, "topological sort:")
	}

	out := []string{}
	for _, n := range nodes {
		out = append(out, g.Vertices[n]...)
	}
	fmt.Println(strings.Join(out, "\n"))
	return nil
}
