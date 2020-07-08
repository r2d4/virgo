package v1

type Graph struct {
	Edges      map[string][]string
	Vertices   map[string][]string
	Directives []Directive
}

type EdgeStmt struct {
	RHS, LHS []string
	Op       string
}

type Directive struct {
	Name string
	Args []KV
}

type KV struct {
	Key   string
	Value string
}

type Edge struct {
	Source, Destination string
}
