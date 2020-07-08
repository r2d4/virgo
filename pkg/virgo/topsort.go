package virgo

import (
	"sort"

	"github.com/pkg/errors"

	v1 "matt-rickard.com/virgo/pkg/types/v1"
)

func DFS(g *v1.Graph, root string) ([]string, error) {
	d := newDFS(g, root)
	return d.sort()
}

func TopSort(g *v1.Graph) ([]string, error) {
	return DFS(g, "")
}

type dfs struct {
	seen, processed map[string]bool
	order           []string
	root            string

	g *v1.Graph
}

func newDFS(g *v1.Graph, root string) *dfs {
	return &dfs{
		g:         g,
		seen:      map[string]bool{},
		processed: map[string]bool{},
		order:     []string{},
		root:      root,
	}
}

func (d *dfs) sort() ([]string, error) {
	var keys []string
	if d.root != "" {
		keys = []string{d.root}
	} else {
		for n := range d.g.Edges {
			keys = append(keys, n)
		}
		sort.StringSlice.Sort(keys)
	}
	for _, n := range keys {
		if !d.processed[n] {
			if err := d.visit(n); err != nil {
				return nil, errors.Wrapf(err, "visiting %s", n)
			}
		}
	}
	return d.order, nil
}

func (d *dfs) visit(n string) error {
	if d.processed[n] {
		return nil
	}
	if d.seen[n] {
		return errors.New("cycle exists, can't sort")
	}
	d.seen[n] = true

	for _, m := range d.g.Edges[n] {
		if err := d.visit(m); err != nil {
			return errors.Wrapf(err, "visiting %s", m)
		}
	}
	d.seen[n] = false
	d.processed[n] = true
	d.order = append([]string{n}, d.order...)
	return nil
}
