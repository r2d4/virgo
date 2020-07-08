package virgo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "matt-rickard.com/virgo/pkg/types/v1"
)

func TestTopSort(t *testing.T) {
	tests := []struct {
		desc string
		in   *v1.Graph
		out  []string
		err  bool
	}{
		{
			desc: "basic",
			in: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b", "d"},
					"b": {"c", "e"},
					"c": {},
					"e": {"c", "f"},
					"d": {"e"},
					"f": {},
					"g": {"d"},
				},
			},
			out: []string{"g", "a", "d", "b", "e", "f", "c"},
		},
		{
			desc: "cycle",
			in:   graphCycle,
			err:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			out, err := TopSort(test.in)
			if err != nil && !test.err {
				t.Errorf("unexpected error: %s", err)
				return
			}
			if test.err && err == nil {
				t.Errorf("expected error but got none: %+v", out)
				return
			}
			if diff := cmp.Diff(test.out, out); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

var graph1 = &v1.Graph{
	Edges: map[string][]string{
		"a": {"b", "d"},
		"b": {"c", "e"},
		"c": {},
		"e": {"c", "f"},
		"d": {"e"},
		"f": {},
		"g": {"d"},
	},
}

var graphTwoLines = &v1.Graph{
	Edges: map[string][]string{
		"root": {"a1", "b1"},
		"a1":   {"a2"},
		"a2":   {"a3"},
		"b1":   {"b2"},
		"b2":   {"b3"},
	},
}

var graphCycle = &v1.Graph{
	Edges: map[string][]string{
		"a": {"b"},
		"b": {"c"},
		"c": {"d"},
		"d": {"b"},
	},
}

func TestDFS(t *testing.T) {
	tests := []struct {
		desc string
		in   *v1.Graph
		root string
		out  []string
		err  bool
	}{
		{
			desc: "basic",
			in:   graph1,
			root: "d",
			out:  []string{"d", "e", "f", "c"},
		},
		{
			desc: "two paths",
			in:   graphTwoLines,
			root: "a2",
			out:  []string{"a2", "a3"},
		},
		{
			desc: "two paths root",
			in:   graphTwoLines,
			root: "root",
			out:  []string{"root", "b1", "b2", "b3", "a1", "a2", "a3"},
		},
		{
			desc: "cycle",
			in:   graphCycle,
			root: "a",
			err:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			out, err := DFS(test.in, test.root)
			if err != nil && !test.err {
				t.Errorf("unexpected error: %s", err)
				return
			}
			if test.err && err == nil {
				t.Errorf("expected error but got none: %+v", out)
				return
			}
			if diff := cmp.Diff(test.out, out); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
