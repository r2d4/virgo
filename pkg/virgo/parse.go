package virgo

import (
	"matt-rickard.com/virgo/pkg/parser"
	v1 "matt-rickard.com/virgo/pkg/types/v1"
)

// Parse parses a virgo configuration file into a v1.Graph
func Parse(input []byte) (*v1.Graph, error) {
	return parser.Parse(input)
}
