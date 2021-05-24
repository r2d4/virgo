package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "matt-rickard.com/virgo/pkg/types/v1"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc, in string
		out      *v1.Graph
		err      bool
	}{
		{
			desc: "basic",
			in: `
			a -> b

			`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "two edges",
			in: `

					a -> b
					b -> c

					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
					"b": {"c"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "comment",
			in: `
					// comment
					a -> b
					b -> c
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
					"b": {"c"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "long edge name",
			in: `
					a_long_edge_name -> an_even_longer_edge_name
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a_long_edge_name": {"an_even_longer_edge_name"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "right and left edge",
			in: `
					// comment
					a -> b
					c <- b
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
					"b": {"c"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "undirected edge",
			in: `
					a -- b
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "edge ending with digit",
			in: `
					a_0 -> a_1
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a_0": {"a_1"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "4-node list edge",
			in: `
					a -> b, c, d, e
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b", "c", "d", "e"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "quoted edges",
			in: `
					"a string" -> "b", c
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a string": {"b", "c"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "variatic node names",
			in: `
					a, b, c -> d, e
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"d", "e"},
					"b": {"d", "e"},
					"c": {"d", "e"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "7 edges one line",
			in: `
					a -> b -> c -> d -> e -> f -> g`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b"},
					"b": {"c"},
					"c": {"d"},
					"d": {"e"},
					"e": {"f"},
					"f": {"g"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "multi-edge and node list",
			in: `
					a -> b,c,d -> e`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"a": {"b", "c", "d"},
					"b": {"e"},
					"c": {"e"},
					"d": {"e"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "left op",
			in:   "a <- b",
			out: &v1.Graph{
				Edges: map[string][]string{
					"b": {"a"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "right and left op together",
			in: `
					dest_1 <- src -> dest_2 -> dest_3
					src_1 -> dest <- src_2
					`,
			out: &v1.Graph{
				Edges: map[string][]string{
					"src":    {"dest_1", "dest_2"},
					"dest_2": {"dest_3"},
					"src_1":  {"dest"},
					"src_2":  {"dest"},
				},
				Vertices: map[string][]string{},
			},
		},
		{
			desc: "vertex definition",
			in: "#directive\n" +
				"#directive_2 arg=value arg2=value2\n" +
				"src -> dest\n" +
				"src = `src value`\n" +
				"dest = ` dest value 1\n" +
				"dest value 2`",
			out: &v1.Graph{
				Directives: []v1.Directive{
					{Name: "directive"},
					{
						Name: "directive_2",
						Args: []v1.KV{
							{Key: "arg", Value: "value"},
							{Key: "arg2", Value: "value2"},
						},
					},
				},
				Edges: map[string][]string{
					"src": {"dest"},
				},
				Vertices: map[string][]string{
					"src":  {"src value"},
					"dest": {"dest value 1", "dest value 2"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			out, err := Parse([]byte(test.in))
			if err != nil && !test.err {
				t.Errorf("unexpected error: %+v", err)
				return
			}
			if err == nil && test.err {
				t.Errorf("expected error, but got none: output: %+v", out)
				return
			}
			if diff := cmp.Diff(test.out, out); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
