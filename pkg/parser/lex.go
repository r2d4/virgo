package parser

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	v1 "matt-rickard.com/virgo/pkg/types/v1"
)

// Parse parses the input and returs the result.
func Parse(input []byte) (*v1.Graph, error) {
	l := newLexer(input)
	l.graph = &v1.Graph{
		Edges:    map[string][]string{},
		Vertices: map[string][]string{},
	}
	_ = yyParse(l)
	return l.graph, l.err
}

func (l *lexer) AddEdges(op string, src, dest []string) {
	if op == "<-" {
		src, dest = dest, src
	}
	for _, s := range src {
		if _, ok := l.graph.Edges[s]; !ok {
			l.graph.Edges[s] = []string{}
		}
		for _, d := range dest {
			var seen bool
			for _, n := range l.graph.Edges[s] {
				if n == d {
					seen = true
					break
				}
			}
			if !seen {
				l.graph.Edges[s] = append(l.graph.Edges[s], d)
			}
		}
	}
}

func (l *lexer) AddVertexDef(name string, block string) {
	block = strings.Trim(block, "\t \n")
	block = strings.Replace(block, "\t", "", -1)
	lines := strings.Split(block, "\n")
	l.graph.Vertices[name] = lines
}

type lexer struct {
	empty   bool
	current byte

	line int
	pos  int

	err error

	buf   *bytes.Buffer
	src   *bufio.Reader
	graph *v1.Graph
}

func (l *lexer) Error(s string) {
	e := fmt.Sprintf("error: line: %+v, char: %+v, token: %s, err: %+v", l.line, l.current, l.token(), s)
	l.err = errors.New(e)
}

func (l *lexer) getc() byte {
	if l.current != 0 {
		l.buf.Write([]byte{l.current})
	}
	l.current = 0
	if b, err := l.src.ReadByte(); err == nil {
		if b == '\n' {
			l.line++
			l.pos = 0
		} else {
			l.pos++
		}
		l.current = b
	} else if err != io.EOF {
		log.Fatal(err)
	}
	return l.current
}

func (l *lexer) token() string {
	return l.buf.String()
}

func newLexer(src []byte) *lexer {
	l := &lexer{
		src:  bufio.NewReader(bytes.NewReader(src)),
		buf:  new(bytes.Buffer),
		pos:  1,
		line: 1,
	}

	l.getc()
	return l
}
